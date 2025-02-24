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

from polyaxon.polypod.compiler.contexts.kubeflow.mpi_job import MPIJobContextsManager
from polyaxon.polypod.compiler.contexts.kubeflow.mx_job import MXJobContextsManager
from polyaxon.polypod.compiler.contexts.kubeflow.pytroch_job import (
    PytorchJobContextsManager,
)
from polyaxon.polypod.compiler.contexts.kubeflow.tf_job import TfJobContextsManager
from polyaxon.polypod.compiler.contexts.kubeflow.xgb_job import (
    XGBoostJobContextsManager,
)
