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

import pytest

from polyaxon.polypod.common.container_resources import sanitize_resources
from polyaxon.utils.test_utils import BaseTestCase


@pytest.mark.polypod_mark
class TestResourceRequirements(BaseTestCase):
    def test_empty_sanitize_resources(self):
        resources = sanitize_resources(None)
        assert resources is None

    def test_empty_dict_sanitize_resources(self):
        resources = sanitize_resources({})
        assert resources is None

    def test_sanitize_resources_containing_ints(self):
        resources = sanitize_resources(
            {"requests": {"cpu": 1, "memory": "100Mi", "nvidia.com/gpu": 1}}
        )
        assert resources.limits is None
        assert resources.requests == {
            "cpu": "1",
            "memory": "100Mi",
            "nvidia.com/gpu": "1",
        }

        resources = sanitize_resources(
            {"limits": {"cpu": "1", "memory": "100Mi", "nvidia.com/gpu": "1"}}
        )
        assert resources.limits == {
            "cpu": "1",
            "memory": "100Mi",
            "nvidia.com/gpu": "1",
        }
        assert resources.requests is None

        resources = sanitize_resources(
            {
                "requests": {"cpu": "1", "memory": "100Mi", "nvidia.com/gpu": "1"},
                "limits": {"cpu": "1", "memory": "100Mi", "nvidia.com/gpu": "1"},
            }
        )
        assert resources.limits == {
            "cpu": "1",
            "memory": "100Mi",
            "nvidia.com/gpu": "1",
        }
        assert resources.requests == {
            "cpu": "1",
            "memory": "100Mi",
            "nvidia.com/gpu": "1",
        }
