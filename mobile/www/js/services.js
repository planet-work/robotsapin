'use strict';

/* Services */
var sapinAppServices = angular.module('sapinApp.services', ['ngResource']);

sapinAppServices.value('version', '0.0.1');
sapinAppServices.value('api_prefix', './api');

sapinAppServices.factory('Token', ['$window',
  function($window) {
    var Token = {};
    Token.get = function() {
      return $window.sessionStorage.token;
    };
    return Token;
  }
]);

sapinAppServices.factory('APIInterceptor', ['$q', '$rootScope', 'Token',
  function($q, $rootScope, Token) {
    return {
      'request': function(config) {
        var t = Token.get()
        if (t) {
          config.headers['Authorization'] = 'Bearer ' + t;
        }
        return config;
      },
      'requestError': function(rejection) {
        // do something on error
        if (canRecover(rejection)) {
          return responseOrNewPromise
        }
        return $q.reject(rejection);
      },
      'response': function(response) {
        // do something on success
        return response;
      },
      'responseError': function(rejection) {
        if (rejection.status === 400) {
          console.log('evBadRequest');
          $rootScope.$broadcast('evBadRequest');
        } else if (rejection.status === 401) {
          console.log('evUnauthorized');
          $rootScope.$broadcast('evUnauthorized');
        } else if (rejection.status === 403) {
          console.log('evForbidden');
          $rootScope.$broadcast('evForbidden');
        } else if (rejection.status === 404) {
          console.log('evNotFound');
          $rootScope.$broadcast('evNotFound');
        } else if (rejection.status === 500) {
          console.log('evInternalServerError');
          $rootScope.$broadcast('evInternalServerError');
        } else if (rejection.status === 502) {
          console.log('evBadGateway');
          $rootScope.$broadcast('evBadGateway');
        }
        return $q.reject(rejection);
      }
    }
  }
]);

sapinAppServices.factory('Database', ['$resource', 'api_prefix',
  function($resource, api_prefix) {
    return $resource(api_prefix + '/database/:databaseId', {}, {
      get: {
        method: 'GET',
        params: {
          databaseId: ''
        },
      },
      'query': {
        method: 'GET',
        url: api_prefix + '/database/',
        isArray: false,
      },
      'update': {
        method: 'PUT',
        params: {
          databaseId: ''
        },
      },

    });
  }
]);

