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

from polyaxon.connections.base import BaseService


class AzureService(BaseService):
    def __init__(self, connection=None, **kwargs):
        super().__init__(connection=connection, **kwargs)
        self._account_name = kwargs.get("account_name") or kwargs.get(
            "AZURE_ACCOUNT_NAME"
        )
        self._account_key = kwargs.get("account_key") or kwargs.get("AZURE_ACCOUNT_KEY")
        self._connection_string = kwargs.get("connection_string") or kwargs.get(
            "AZURE_CONNECTION_STRING"
        )

    def set_connection(
        self,
        connection=None,
        connection_name=None,
        account_name=None,
        account_key=None,
        connection_string=None,
    ):
        raise NotImplementedError

    def set_env_vars(self):
        if self._account_name:
            os.environ["AZURE_ACCOUNT_NAME"] = self._account_name
        if self._account_key:
            os.environ["AZURE_ACCOUNT_KEY"] = self._account_key
        if self._connection_string:
            os.environ["AZURE_CONNECTION_STRING"] = self._connection_string
