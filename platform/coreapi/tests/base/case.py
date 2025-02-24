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

import tempfile

from django.conf import settings

from coredb.factories.users import UserFactory
from polycommon.test_cases.base import PolyaxonBaseTest
from polycommon.test_clients.base import BaseClient


class BaseTest(PolyaxonBaseTest):
    def setUp(self):
        super().setUp()

        settings.ARTIFACTS_ROOT = tempfile.mkdtemp()
        settings.LOGS_ROOT = tempfile.mkdtemp()
        settings.ARCHIVES_ROOT = tempfile.mkdtemp()
        self.client = BaseClient()
        self.user = UserFactory()
