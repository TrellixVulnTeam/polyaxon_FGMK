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

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.PolyaxonSdk);
  }
}(this, function(expect, PolyaxonSdk) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new PolyaxonSdk.V1RunSchema();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('V1RunSchema', function() {
    it('should create an instance of V1RunSchema', function() {
      // uncomment below and update the code to test V1RunSchema
      //var instane = new PolyaxonSdk.V1RunSchema();
      //expect(instance).to.be.a(PolyaxonSdk.V1RunSchema);
    });

    it('should have the property job (base name: "job")', function() {
      // uncomment below and update the code to test the property job
      //var instane = new PolyaxonSdk.V1RunSchema();
      //expect(instance).to.be();
    });

    it('should have the property service (base name: "service")', function() {
      // uncomment below and update the code to test the property service
      //var instane = new PolyaxonSdk.V1RunSchema();
      //expect(instance).to.be();
    });

    it('should have the property dag (base name: "dag")', function() {
      // uncomment below and update the code to test the property dag
      //var instane = new PolyaxonSdk.V1RunSchema();
      //expect(instance).to.be();
    });

    it('should have the property tfJob (base name: "tfJob")', function() {
      // uncomment below and update the code to test the property tfJob
      //var instane = new PolyaxonSdk.V1RunSchema();
      //expect(instance).to.be();
    });

    it('should have the property pytorchJob (base name: "pytorchJob")', function() {
      // uncomment below and update the code to test the property pytorchJob
      //var instane = new PolyaxonSdk.V1RunSchema();
      //expect(instance).to.be();
    });

    it('should have the property mpiJob (base name: "mpiJob")', function() {
      // uncomment below and update the code to test the property mpiJob
      //var instane = new PolyaxonSdk.V1RunSchema();
      //expect(instance).to.be();
    });

    it('should have the property dask (base name: "dask")', function() {
      // uncomment below and update the code to test the property dask
      //var instane = new PolyaxonSdk.V1RunSchema();
      //expect(instance).to.be();
    });

    it('should have the property spark (base name: "spark")', function() {
      // uncomment below and update the code to test the property spark
      //var instane = new PolyaxonSdk.V1RunSchema();
      //expect(instance).to.be();
    });

    it('should have the property flink (base name: "flink")', function() {
      // uncomment below and update the code to test the property flink
      //var instane = new PolyaxonSdk.V1RunSchema();
      //expect(instance).to.be();
    });

    it('should have the property ruy (base name: "ruy")', function() {
      // uncomment below and update the code to test the property ruy
      //var instane = new PolyaxonSdk.V1RunSchema();
      //expect(instance).to.be();
    });

  });

}));
