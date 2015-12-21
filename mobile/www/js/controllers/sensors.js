'use strict';

angular.module('sapinApp.controllers').controller('SensorsCtrl', function($scope, $stateParams, $filter, $translate, Sensors,$rootScope,$http) {

	var sensorsctrl = this;
    sensorsctrl.sequences =[];

    this.GetSensors = function() {
        Sensors.get({},
  function success(result) {
      sensorsctrl.sequences = result.data;
      $scope.sequences = result.data;
    }, function failure(httpResponse) {
      $scope.erf = "mouarf";
      console.log("err: " + httpResponse);
    });


	};	

	this.GetSensors();
});
