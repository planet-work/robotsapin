'use strict';

/* Services */
var sapinAppServices = angular.module('sapinApp.services', ['ngResource']);

sapinAppServices.value('version', '0.0.1');
sapinAppServices.value('api_prefix', api_prefix);

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
		if (response.config.method != 'GET') {
		    $rootScope.$broadcast('updateStatus');
		}
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


sapinAppServices.factory('Status', ['$resource', 'api_prefix',
  function($resource, api_prefix) {
    return $resource(api_prefix + '/status', {}, {
      get: {
        method: 'GET',
      },
    });
  }
]);



sapinAppServices.factory('Music', ['$resource', 'api_prefix','$rootScope',
  function($resource, api_prefix,$rootScope) {
    return $resource(api_prefix + '/music/:filename', {}, {
      get: {
        method: 'GET',
        params: {
          filename: ''
        },
      },
      'play': {
		method: 'POST',
		url: api_prefix + '/music/',
        params: {
          filename: ''
        },
		isArray: false,
      },
      'pause': {
        method: 'PUT',
        url: api_prefix + '/music/pause',
        params: {
          filename: ''
        },
      },
      'stop': {
        method: 'PUT',
        url: api_prefix + '/music/stop',
        params: {
          filename: ''
        },
      },
      'volup': {
        method: 'PUT',
        url: api_prefix + '/music/volume+',
        params: {
          filename: ''
        },
      },
      'voldown': {
        method: 'PUT',
        url: api_prefix + '/music/volume-',
        params: {
          filename: ''
        },
      },

    });
  }
]);

sapinAppServices.factory('Display', ['$resource', 'api_prefix','$rootScope',
  function($resource, api_prefix,$rootScope) {
    return $resource(api_prefix + '/display/:filename', {}, {
      get: {
        method: 'GET',
        params: {
          filename: ''
        },
      },
      'show': {
		method: 'POST',
		url: api_prefix + '/display/',
        params: {
          filename: ''
        },
		isArray: false,
      },
      'clear': {
        method: 'PUT',
        url: api_prefix + '/display/clear',
        params: {
          filename: ''
        },
      }
    });
  }
]);

sapinAppServices.factory('Topper', ['$resource', 'api_prefix','$rootScope',
  function($resource, api_prefix,$rootScope) {
    return $resource(api_prefix + '/topper/:filename', {}, {
      get: {
        method: 'GET',
        params: {
          filename: ''
        },
      },
      'start': {
		method: 'POST',
		url: api_prefix + '/topper/',
        params: {
          filename: ''
        },
		isArray: false,
      },
      'speed': {
		method: 'PUT',
		url: api_prefix + '/topper/speed',
        params: {
          speed: ''
        },
		isArray: false,
      }
    });
  }
]);
