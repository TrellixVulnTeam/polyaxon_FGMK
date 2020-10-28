// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.2.1-rc5
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1HyperoptAlgorithms from './V1HyperoptAlgorithms';

/**
 * The V1Hyperopt model module.
 * @module model/V1Hyperopt
 * @version 1.2.1-rc5
 */
class V1Hyperopt {
    /**
     * Constructs a new <code>V1Hyperopt</code>.
     * @alias module:model/V1Hyperopt
     */
    constructor() { 
        
        V1Hyperopt.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1Hyperopt</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Hyperopt} obj Optional instance to populate.
     * @return {module:model/V1Hyperopt} The populated <code>V1Hyperopt</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Hyperopt();

            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
            if (data.hasOwnProperty('algorithm')) {
                obj['algorithm'] = V1HyperoptAlgorithms.constructFromObject(data['algorithm']);
            }
            if (data.hasOwnProperty('params')) {
                obj['params'] = ApiClient.convertToType(data['params'], {'String': Object});
            }
            if (data.hasOwnProperty('num_runs')) {
                obj['num_runs'] = ApiClient.convertToType(data['num_runs'], 'Number');
            }
            if (data.hasOwnProperty('max_iterations')) {
                obj['max_iterations'] = ApiClient.convertToType(data['max_iterations'], 'Number');
            }
            if (data.hasOwnProperty('seed')) {
                obj['seed'] = ApiClient.convertToType(data['seed'], 'Number');
            }
            if (data.hasOwnProperty('concurrency')) {
                obj['concurrency'] = ApiClient.convertToType(data['concurrency'], 'Number');
            }
            if (data.hasOwnProperty('early_stopping')) {
                obj['early_stopping'] = ApiClient.convertToType(data['early_stopping'], [Object]);
            }
        }
        return obj;
    }


}

/**
 * @member {String} kind
 * @default 'hyperopt'
 */
V1Hyperopt.prototype['kind'] = 'hyperopt';

/**
 * @member {module:model/V1HyperoptAlgorithms} algorithm
 */
V1Hyperopt.prototype['algorithm'] = undefined;

/**
 * @member {Object.<String, Object>} params
 */
V1Hyperopt.prototype['params'] = undefined;

/**
 * @member {Number} num_runs
 */
V1Hyperopt.prototype['num_runs'] = undefined;

/**
 * @member {Number} max_iterations
 */
V1Hyperopt.prototype['max_iterations'] = undefined;

/**
 * @member {Number} seed
 */
V1Hyperopt.prototype['seed'] = undefined;

/**
 * @member {Number} concurrency
 */
V1Hyperopt.prototype['concurrency'] = undefined;

/**
 * @member {Array.<Object>} early_stopping
 */
V1Hyperopt.prototype['early_stopping'] = undefined;






export default V1Hyperopt;

