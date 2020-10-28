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
    instance = new PolyaxonSdk.ProjectsV1Api();
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

  describe('ProjectsV1Api', function() {
    describe('archiveProject', function() {
      it('should call archiveProject successfully', function(done) {
        //uncomment below and update the code to test archiveProject
        //instance.archiveProject(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('bookmarkProject', function() {
      it('should call bookmarkProject successfully', function(done) {
        //uncomment below and update the code to test bookmarkProject
        //instance.bookmarkProject(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('createProject', function() {
      it('should call createProject successfully', function(done) {
        //uncomment below and update the code to test createProject
        //instance.createProject(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteProject', function() {
      it('should call deleteProject successfully', function(done) {
        //uncomment below and update the code to test deleteProject
        //instance.deleteProject(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('disableProjectCI', function() {
      it('should call disableProjectCI successfully', function(done) {
        //uncomment below and update the code to test disableProjectCI
        //instance.disableProjectCI(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('enableProjectCI', function() {
      it('should call enableProjectCI successfully', function(done) {
        //uncomment below and update the code to test enableProjectCI
        //instance.enableProjectCI(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('fetchProjectTeams', function() {
      it('should call fetchProjectTeams successfully', function(done) {
        //uncomment below and update the code to test fetchProjectTeams
        //instance.fetchProjectTeams(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getProject', function() {
      it('should call getProject successfully', function(done) {
        //uncomment below and update the code to test getProject
        //instance.getProject(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getProjectSettings', function() {
      it('should call getProjectSettings successfully', function(done) {
        //uncomment below and update the code to test getProjectSettings
        //instance.getProjectSettings(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listArchivedProjects', function() {
      it('should call listArchivedProjects successfully', function(done) {
        //uncomment below and update the code to test listArchivedProjects
        //instance.listArchivedProjects(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listBookmarkedProjects', function() {
      it('should call listBookmarkedProjects successfully', function(done) {
        //uncomment below and update the code to test listBookmarkedProjects
        //instance.listBookmarkedProjects(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listProjectNames', function() {
      it('should call listProjectNames successfully', function(done) {
        //uncomment below and update the code to test listProjectNames
        //instance.listProjectNames(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listProjects', function() {
      it('should call listProjects successfully', function(done) {
        //uncomment below and update the code to test listProjects
        //instance.listProjects(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('patchProject', function() {
      it('should call patchProject successfully', function(done) {
        //uncomment below and update the code to test patchProject
        //instance.patchProject(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('patchProjectSettings', function() {
      it('should call patchProjectSettings successfully', function(done) {
        //uncomment below and update the code to test patchProjectSettings
        //instance.patchProjectSettings(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('patchProjectTeams', function() {
      it('should call patchProjectTeams successfully', function(done) {
        //uncomment below and update the code to test patchProjectTeams
        //instance.patchProjectTeams(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('restoreProject', function() {
      it('should call restoreProject successfully', function(done) {
        //uncomment below and update the code to test restoreProject
        //instance.restoreProject(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('unbookmarkProject', function() {
      it('should call unbookmarkProject successfully', function(done) {
        //uncomment below and update the code to test unbookmarkProject
        //instance.unbookmarkProject(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateProject', function() {
      it('should call updateProject successfully', function(done) {
        //uncomment below and update the code to test updateProject
        //instance.updateProject(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateProjectSettings', function() {
      it('should call updateProjectSettings successfully', function(done) {
        //uncomment below and update the code to test updateProjectSettings
        //instance.updateProjectSettings(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateProjectTeams', function() {
      it('should call updateProjectTeams successfully', function(done) {
        //uncomment below and update the code to test updateProjectTeams
        //instance.updateProjectTeams(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('uploadProjectArtifact', function() {
      it('should call uploadProjectArtifact successfully', function(done) {
        //uncomment below and update the code to test uploadProjectArtifact
        //instance.uploadProjectArtifact(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
