'use strict';

angular.module('sapinApp.controllers').controller('DisplayCtrl', function($scope, $stateParams, $filter, $translate, Display,$rootScope,$http) {

	var displayctrl = this;
    displayctrl.images =[];

    this.ListDisplay = function() {
        Display.get({},
  function success(result) {
      displayctrl.songs = result.data;
      $scope.images = result.data;
    }, function failure(httpResponse) {
      $scope.erf = "mouarf";
      console.log("err: " + httpResponse);
    });


	};	

	this.ListDisplay();

	$scope.Clear = function() {
        Display.clear();
	};

	$scope.Show = function(image) {
		console.log("Display image " + image);
        $http.post(api_prefix + '/display/' + image);
		//$rootScope.filename = song;
        //Display.play({filename: song});
		//$rootScope.$broadcast('updateStatus');
	};
});
