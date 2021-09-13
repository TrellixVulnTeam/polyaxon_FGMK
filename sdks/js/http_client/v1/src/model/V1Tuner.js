// Copyright 2018-2021 Polyaxon, Inc.
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
 * The version of the OpenAPI document: 1.11.2
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1Param from './V1Param';

/**
 * The V1Tuner model module.
 * @module model/V1Tuner
 * @version 1.11.2
 */
class V1Tuner {
    /**
     * Constructs a new <code>V1Tuner</code>.
     * @alias module:model/V1Tuner
     */
    constructor() { 
        
        V1Tuner.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1Tuner</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Tuner} obj Optional instance to populate.
     * @return {module:model/V1Tuner} The populated <code>V1Tuner</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Tuner();

            if (data.hasOwnProperty('hubRef')) {
                obj['hubRef'] = ApiClient.convertToType(data['hubRef'], 'String');
            }
            if (data.hasOwnProperty('queue')) {
                obj['queue'] = ApiClient.convertToType(data['queue'], 'String');
            }
            if (data.hasOwnProperty('presets')) {
                obj['presets'] = ApiClient.convertToType(data['presets'], ['String']);
            }
            if (data.hasOwnProperty('params')) {
                obj['params'] = ApiClient.convertToType(data['params'], {'String': V1Param});
            }
        }
        return obj;
    }


}

/**
 * @member {String} hubRef
 */
V1Tuner.prototype['hubRef'] = undefined;

/**
 * @member {String} queue
 */
V1Tuner.prototype['queue'] = undefined;

/**
 * @member {Array.<String>} presets
 */
V1Tuner.prototype['presets'] = undefined;

/**
 * @member {Object.<String, module:model/V1Param>} params
 */
V1Tuner.prototype['params'] = undefined;






export default V1Tuner;

