'use strict';

angular.module('sapinApp.controllers').controller('TopperCtrl', function($scope, $stateParams, $filter, $translate, Topper,$rootScope,$http) {

	var topperctrl = this;
    topperctrl.sequences =[];

    this.ListTopper = function() {
        Topper.get({},
  function success(result) {
      topperctrl.sequences = result.data;
      $scope.sequences = result.data;
    }, function failure(httpResponse) {
      $scope.erf = "mouarf";
      console.log("err: " + httpResponse);
    });


	};	

	this.ListTopper();

	$scope.Show = function(image) {
		console.log("Topper LED sequence " + seqId);
        $http.post(api_prefix + '/topper/' + image);
		//$rootScope.filename = song;
        //Topper.play({filename: song});
		//$rootScope.$broadcast('updateStatus');
	};
});
