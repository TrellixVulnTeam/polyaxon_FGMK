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

from polyaxon.managers.agent import AgentConfigManager, SandboxConfigManager
from polyaxon.schemas.cli.agent_config import AgentConfig, SandboxConfig
from polyaxon.utils.test_utils import BaseTestCase


@pytest.mark.managers_mark
class TestAgentConfigManager(BaseTestCase):
    def test_default_props(self):
        assert AgentConfigManager.is_global() is True
        assert AgentConfigManager.CONFIG_PATH is None
        assert AgentConfigManager.IS_POLYAXON_DIR is False
        assert AgentConfigManager.CONFIG_FILE_NAME == ".agent"
        assert AgentConfigManager.CONFIG == AgentConfig


@pytest.mark.managers_mark
class TestSandboxConfigManager(BaseTestCase):
    def test_default_props(self):
        assert SandboxConfigManager.is_global() is False
        assert SandboxConfigManager.is_all_visibility() is True
        assert SandboxConfigManager.CONFIG_PATH is None
        assert SandboxConfigManager.IS_POLYAXON_DIR is False
        assert SandboxConfigManager.CONFIG_FILE_NAME == ".sandbox"
        assert SandboxConfigManager.CONFIG == SandboxConfig
