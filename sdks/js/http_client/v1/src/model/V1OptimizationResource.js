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
import V1ResourceType from './V1ResourceType';

/**
 * The V1OptimizationResource model module.
 * @module model/V1OptimizationResource
 * @version 1.2.1-rc5
 */
class V1OptimizationResource {
    /**
     * Constructs a new <code>V1OptimizationResource</code>.
     * @alias module:model/V1OptimizationResource
     */
    constructor() { 
        
        V1OptimizationResource.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1OptimizationResource</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1OptimizationResource} obj Optional instance to populate.
     * @return {module:model/V1OptimizationResource} The populated <code>V1OptimizationResource</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1OptimizationResource();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = V1ResourceType.constructFromObject(data['type']);
            }
        }
        return obj;
    }


}

/**
 * @member {String} name
 */
V1OptimizationResource.prototype['name'] = undefined;

/**
 * @member {module:model/V1ResourceType} type
 */
V1OptimizationResource.prototype['type'] = undefined;






export default V1OptimizationResource;

