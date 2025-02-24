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
import os

import click

from polyaxon.exceptions import PolyaxonException


@click.command()
@click.argument("component")
@click.option(
    "--path", default="./web", help="Path where the config should be generated."
)
@click.option("--root", help="Absolute root where the configs, default to pwd")
def proxy(component, path, root):
    """Create api proxy."""
    from polyaxon.proxies.generators import (
        generate_api_conf,
        generate_forward_proxy_cmd,
        generate_gateway_conf,
        generate_streams_conf,
    )
    from polyaxon.services.values import PolyaxonServices
    from polyaxon.settings import set_proxies_config

    if not root:
        root = os.path.abspath(".")

    set_proxies_config()

    if PolyaxonServices.is_api(component):
        generate_api_conf(path=path, root=root)
    elif PolyaxonServices.is_streams(component):
        generate_streams_conf(path=path, root=root)
    elif PolyaxonServices.is_gateway(component):
        generate_gateway_conf(path=path, root=root)
        generate_forward_proxy_cmd(path=path)
    else:
        raise PolyaxonException("Component {} is not recognized".format(component))
