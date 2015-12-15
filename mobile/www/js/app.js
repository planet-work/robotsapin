'use strict';

angular.module('fui',['fui.radio']);
// Declare app level module which depends on filters, and servicesi
angular.module('sapinApp', [
    'ngRoute',
    'ngMessages',
    'xeditable',
    'ui.bootstrap',
    'ui.router',
    'pascalprecht.translate',
    'ngTable',
    'truncate',
    'sapinApp.filters',
    'sapinApp.services',
    'sapinApp.directives',
    'sapinApp.controllers',
    'fui',
  ])
  .run(
    ['$rootScope', '$state', '$stateParams', 'editableOptions',
      function($rootScope, $state, $stateParams, editableOptions) {
        // It's very handy to add references to $state and $stateParams to the $rootScope
        // so that you can access them from any scope within your applications.For example,
        // <li ng-class="{ active: $state.includes('contacts.list') }"> will set the <li>
        // to active whenever 'contacts.list' or one of its decendents is active.
        $rootScope.$state = $state;
        $rootScope.$stateParams = $stateParams;
        $rootScope.$on('evUnauthorized', function(event) {
          $rootScope.loggedIn = false;
          $state.go('login');
        });
        $rootScope.$on('evLogin', function(event) {
          $rootScope.loggedIn = true;
          $state.go('dashboard');
        });
        $rootScope.$on('evLogout', function(event) {
          $rootScope.loggedIn = false;
          $state.go('login');
        });
        /*$rootScope.$on('evNotFound', function(event) {
          $state.go('notfound');
        });*/
        editableOptions.theme = 'bs3';
      }
    ]
  )
  .config(['$resourceProvider', '$stateProvider', '$urlRouterProvider', '$httpProvider', '$translateProvider',
    function($resourceProvider, $stateProvider, $urlRouterProvider, $httpProvider, $translateProvider) {
      // State base router
      $resourceProvider.defaults.stripTrailingSlashes = false;
      $urlRouterProvider.otherwise("/status");
      $stateProvider
        .state('status', {
          url: "/status",
          views: {
            'sidebar': {
              templateUrl: 'partials/sidebar.html',
            },
            'content': {
              templateUrl: 'partials/status.html',
            }
          }

        })
        .state('notfound', {
          url: "/notfound",
          templateUrl: "partials/notfound.html"
        })
        .state('dashboard', {
          url: "/dashboard",
          views: {
            'sidebar': {
              templateUrl: 'partials/sidebar.html',
            },
            'content': {
              templateUrl: 'partials/dashboard.html',
              controller: 'DashboardCtrl'
            }
          }
        })
        .state('database.mysql.detail', {
          url: "/{mysqldbId}",
          views: {
            'sidebar@database': {},
            'content@database': {
              templateUrl: 'partials/database/mysql/detail.html',
              controller: 'DatabaseMysqlDetailCtrl',
            }
          }
        });
      // Interceptor
      $httpProvider.interceptors.push('APIInterceptor');
      // Translator
      $translateProvider
        .useStaticFilesLoader({
          prefix: 'locales/',
          suffix: '.json'
        })
        .preferredLanguage('fr')
        .useSanitizeValueStrategy('escape');
    }
  ]);

