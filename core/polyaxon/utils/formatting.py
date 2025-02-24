#!/usr/bin/python
#
# Copyright 2018-2022 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import json
import sys
import urllib.parse as urlparse
import yaml

from typing import Dict, List, Tuple
from urllib.parse import parse_qs

import click

from rich import box
from rich.console import Console
from rich.live import Live
from rich.markdown import Markdown
from rich.progress import (
    BarColumn,
    DownloadColumn,
    Progress,
    TaskProgressColumn,
    TextColumn,
    TimeElapsedColumn,
    TimeRemainingColumn,
    TransferSpeedColumn,
)
from rich.syntax import Syntax
from rich.table import Column, Table
from rich.theme import Theme

from polyaxon.schemas.api.resources import ContainerResourcesConfig
from polyaxon.utils.humanize import humanize_timesince
from polyaxon.utils.list_utils import to_list
from traceml.processors.units_processors import to_percentage, to_unit_memory


def get_meta_response(response):
    def get_pagination(url):
        parsed = urlparse.urlparse(url)
        parsed_query = parse_qs(parsed.query)
        limit = parsed_query.get("limit")
        offset = parsed_query.get("offset")
        res = []
        if limit is not None:
            res.append("--limit={}".format(limit[0]))
        if offset is not None:
            res.append("--offset={}".format(offset[0]))
        return " ".join(res)

    results = {}
    if response.next:
        try:
            results["next"] = get_pagination(response.next)
        except Exception:  # noqa
            results["next"] = response.next
    if response.previous:
        try:
            results["previous"] = get_pagination(response.previous)
        except Exception:  # noqa
            results["previous"] = response.previous
    if response.count:
        results["count"] = response.count
    return results


def humanize_attrs(key, value, rounding=2):
    if key in [
        "created_at",
        "updated_at",
        "started_at",
        "finished_at",
        "schedule_at",
        "last_update_time",
        "last_transition_time",
    ]:
        return humanize_timesince(value)
    if key in ["cpu_percentage"]:
        return to_percentage(value, rounding)
    if key in ["memory_free", "memory_used", "memory_total"]:
        return to_unit_memory(value)
    if isinstance(value, str):
        return value
    if value is None:
        return ""
    return json.dumps(value)


def list_dicts_to_tabulate(
    list_dicts,
    exclude_attrs=None,
    include_attrs=None,
    humanize_values=True,
    upper_keys: bool = True,
):
    exclude_attrs = exclude_attrs or {}
    results = []
    if include_attrs:  # If include_attrs disable exclude_attrs
        exclude_attrs = {}
    for d_value in list_dicts:
        r_value = {}
        for k, v in d_value.items():
            if k in exclude_attrs:
                continue
            if include_attrs and k not in include_attrs:
                continue

            if humanize_values:
                v = humanize_attrs(k, v)

            if upper_keys:
                k = k.upper()
            r_value[k] = v
        results.append(r_value)
    return results


def list_dicts_to_csv(
    list_dicts,
    exclude_attrs=None,
    include_attrs=None,
):
    exclude_attrs = exclude_attrs or {}
    results = []
    if include_attrs:  # If include_attrs disable exclude_attrs
        exclude_attrs = {}
    for d_value in list_dicts:
        result = {}
        for k, v in d_value.items():
            if k in exclude_attrs:
                continue
            if include_attrs and k not in include_attrs:
                continue
            result[k] = v
        results.append(result)

    return results


def dict_to_tabulate(d_value, exclude_attrs=None, humanize_values=True):
    exclude_attrs = exclude_attrs or {}
    results = {}
    for k, v in d_value.items():
        if k in exclude_attrs:
            continue

        if humanize_values:
            v = humanize_attrs(k, v)

        results[k.upper()] = v

    return results


def dict_tabulate(dict_value, is_list_dict=False):
    if is_list_dict:
        headers = dict_value[0].keys() if dict_value else []
        table = Printer.get_table(*headers)
        for d in dict_value:
            table.add_row(*d.values())
        Printer.print(table)
    else:
        table = Printer.get_table(show_header=False, padding=0, box=box.SIMPLE)
        for k, v in dict_value.items():
            table.add_row(k, humanize_attrs(k, v))
        Printer.print(table)


class Printer:
    COLORS = ["yellow", "blue", "magenta", "green", "cyan", "red", "white"]
    console = Console(
        theme=Theme(
            {
                "header": "yellow",
                "success": "green",
                "info": "cyan",
                "warning": "magenta",
                "error": "red",
                "white": "white",
            }
        ),
        markup=True,
    )

    @staticmethod
    def get_progress():
        return Progress(
            TextColumn("[progress.description]{task.description}"),
            BarColumn(),
            DownloadColumn(),
            TransferSpeedColumn(),
            TaskProgressColumn(),
            TextColumn("eta"),
            TimeRemainingColumn(),
            TextColumn("elapsed"),
            TimeElapsedColumn(),
        )

    @classmethod
    def get_live(cls):
        return Live(console=cls.console)

    @staticmethod
    def get_table(*args, **kwargs):
        return Table(*[Column(header=h, no_wrap=True) for h in args], **kwargs)

    @staticmethod
    def pprint(value):
        """Prints as formatted JSON"""
        click.echo(json.dumps(value, sort_keys=True, indent=4, separators=(",", ": ")))

    @classmethod
    def print_md(cls, md: str):
        cls.console.print(Markdown(md))

    @classmethod
    def print_text(cls, value: str):
        syntax = Syntax(value, "txt", theme="dracula", line_numbers=False)
        cls.console.print(syntax)

    @classmethod
    def print_yaml(cls, value: any):
        if isinstance(value, str):
            value = yaml.safe_load(value)
        value = yaml.safe_dump(value, sort_keys=True, indent=2)
        syntax = Syntax(value, "yaml", theme="dracula", line_numbers=False)
        cls.console.print(syntax)

    @classmethod
    def print_json(cls, value: any):
        if isinstance(value, str):
            value = json.loads(value)
        value = json.dumps(value, sort_keys=True, indent=4, separators=(",", ": "))
        syntax = Syntax(value, "json", theme="dracula", line_numbers=False)
        cls.console.print(syntax)

    @classmethod
    def print_help(cls, command_help: str = None):
        if command_help:
            cls.console.print(
                "Please run [white]`polyaxon {} --help`[/white] for more details".format(
                    command_help
                ),
                style="info",
            )

    @classmethod
    def print(cls, text):
        cls.console.print(text)

    @classmethod
    def print_heading(cls, text):
        cls.print_header("\n{}\n".format(text))

    @classmethod
    def print_header(cls, text):
        cls.console.print(text, style="header")

    @classmethod
    def print_warning(cls, text, command_help: str = None):
        cls.console.print(text, style="warning")
        if command_help:
            cls.print_help(command_help)

    @classmethod
    def print_success(cls, text):
        cls.console.print(text, style="success")

    @classmethod
    def print_error(cls, text, sys_exit: bool = False, command_help: str = None):
        cls.console.print(text, style="error")
        if command_help:
            cls.print_help(command_help)
        if sys_exit:
            sys.exit(1)

    @classmethod
    def print_tip(cls, text):
        cls.console.print(text, style="white")

    @classmethod
    def print_info(cls, text):
        cls.console.print(text, style="info")

    @staticmethod
    def add_log_color(value, color):
        return click.style("{}".format(value), fg=color)

    @classmethod
    def add_color(cls, value, style):
        return "[{style}]{value}[/{style}]".format(value=value, style=style)

    @classmethod
    def get_colored_status(cls, status):
        if status == "created":
            return cls.add_color(status, "info")
        elif status == "succeeded":
            return cls.add_color(status, style="success")
        elif status in ["failed", "stopped", "upstream_failed"]:
            return cls.add_color(status, style="error")
        elif status == "done":
            return cls.add_color(status, style="white")

        return cls.add_color(status, style="header")

    @classmethod
    def add_status_color(cls, obj_dict, status_key="status"):
        if obj_dict.get(status_key) is None:
            return obj_dict

        obj_dict[status_key] = cls.get_colored_status(obj_dict[status_key])
        return obj_dict

    @classmethod
    def add_memory_unit(cls, obj_dict, keys):
        keys = to_list(keys)
        for key in keys:
            obj_dict[key] = to_unit_memory(obj_dict[key])
        return obj_dict

    @classmethod
    def decorate_format_value(cls, text_format, values, color):
        values = to_list(values)
        values = [cls.add_color(value, color) for value in values]
        click.echo(text_format.format(*values))

    @staticmethod
    def log(value, nl=False):
        click.echo(value, nl=nl)

    @classmethod
    def resources(cls, jobs_resources):
        jobs_resources = to_list(jobs_resources)
        click.clear()
        table = Printer.get_table("Job", "Mem Usage / Total", "CPU% - CPUs")
        for job_resources in jobs_resources:
            job_resources = ContainerResourcesConfig.from_dict(job_resources)
            line = [
                job_resources.job_name,
                "{} / {}".format(
                    to_unit_memory(job_resources.memory_used),
                    to_unit_memory(job_resources.memory_limit),
                ),
                "{} - {}".format(
                    to_percentage(job_resources.cpu_percentage / 100),
                    job_resources.n_cpus,
                ),
            ]
            table.add_row(*line)
        Printer.print(table)
        sys.stdout.flush()

    @classmethod
    def gpu_resources(cls, jobs_resources):
        jobs_resources = to_list(jobs_resources)
        click.clear()
        table = Printer.get_table(
            "job_name",
            "name",
            "GPU Usage",
            "GPU Mem Usage / Total",
            "GPU Temperature",
            "Power Draw / Limit",
        )
        non_gpu_jobs = 0
        for job_resources in jobs_resources:
            job_resources = ContainerResourcesConfig.from_dict(job_resources)
            line = []
            if not job_resources.gpu_resources:
                non_gpu_jobs += 1
                continue
            for gpu_resources in job_resources.gpu_resources:
                line += [
                    job_resources.job_name,
                    gpu_resources.name,
                    to_percentage(gpu_resources.utilization_gpu / 100),
                    "{} / {}".format(
                        to_unit_memory(gpu_resources.memory_used),
                        to_unit_memory(gpu_resources.memory_total),
                    ),
                    gpu_resources.temperature_gpu,
                    "{} / {}".format(
                        gpu_resources.power_draw, gpu_resources.power_limit
                    ),
                ]
            table.add_row(*line)
        if non_gpu_jobs == len(jobs_resources):
            Printer.print_error(
                "No GPU job was found, please run `resources` command without `-g | --gpu` option."
            )
            exit(1)
        Printer.print(table)
        sys.stdout.flush()


def flatten_keys(
    objects: List[Dict], columns: List[str], columns_prefix: Dict = None
) -> Tuple[List[Dict], Dict]:
    # Extend run with params_keys
    keys = set([])
    columns_prefix = columns_prefix or {}
    prefixed_columns = {}

    def process_objects():
        results = {}
        for k, v in col_values.items():
            results["{}.{}".format(column_prefix, k)] = v
            prefixed_columns[k] = "{}.{}".format(column_prefix, k)
        return results

    for col in columns:
        for obj in objects:
            col_values = obj.pop(col, {}) or {}
            if col in columns_prefix:
                column_prefix = columns_prefix[col]
                col_values = process_objects()
            col_keys = col_values.keys()
            keys |= set(col_keys)
            obj.update(col_values)

    # Check that all obj have all metrics
    # TODO: optimize this process
    for obj in objects:
        obj_keys = set(obj.keys())
        for key in keys:
            if key not in obj_keys:
                obj[key] = None

    return objects, prefixed_columns
