'use strict';

angular.module('sapinApp.controllers')
  .controller('AccountCtrl', ['$scope',
    function($scope) {

    }
  ])

.controller('NavBarCtrl', ['$scope', 'UserData',
  function($scope, UserData) {
    $scope.contact_id = UserData.getId();
  }
])

.controller('MainCtrl', ['$rootScope', '$scope',
  function($rootScope, $scope) {

    $scope.alerts = [];

    $rootScope.$on('evBadGateway', function(event) {
      $scope.addAlert(event);
    });

    $rootScope.$on('evInternalServerError', function(event) {
      $scope.addAlert(event);
    });

    $rootScope.$on('evNotFound', function(event) {
      $scope.addAlert(event);
    });

    /*$rootScope.$on('evForbidden', function(event) {
      $scope.addAlert(event);
    });*/

    $scope.addAlert = function(event) {
      var skip;
      angular.forEach($scope.alerts, function(alert) {
        if (alert.msg == event.name) {
          skip = true;
        }
      });
      if (!skip) {
        $scope.alerts.push({
          type: 'danger',
          msg: event.name
        })
      }
    };

    $scope.closeAlert = function(index) {
      $scope.alerts.splice(index, 1);
    };

  }
]);

