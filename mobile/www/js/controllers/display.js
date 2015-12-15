use strict';

angular.module('sapinApp.controllers').controller('DatabaseAggregateCtrl', function($scope, $stateParams, $filter, $translate, NgTableParams, TableTricks, Database, MysqlDb) {

  $scope.addCollapsed = true;
  $scope.newDatabase = {
    type: 'person',
    lang: 'fr'
  }
  $scope.langs = ['fr'];
  $scope.navPills = [{
    "name": "mysqldbs",
    "active": false
  }]
});
