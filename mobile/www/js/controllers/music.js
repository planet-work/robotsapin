'use strict';

angular.module('sapinApp.controllers').controller('MusicCtrl', function($scope, $stateParams, $filter, $translate, Music) {

	var musicctrl = this;
    musicctrl.songs =[];

    this.ListMusic = function() {
        Music.get({},
  function success(result) {
      musicctrl.songs = result.data;
      $scope.songs = result.data;
    }, function failure(httpResponse) {
      $scope.erf = "mouarf";
      console.log("err: " + httpResponse);
    });


	};	

	this.ListMusic();

	$scope.Stop = function() {
        Music.stop();
	};

	$scope.Pause = function() {
        Music.pause();
	};

	$scope.VolUp = function() {
        Music.volup();
	};

	$scope.VolDown = function() {
        Music.voldown();
	};
    
	$scope.Play = function(song) {
		console.log("Play music " + song);
		Music.play({"filename": song});
	};
});
