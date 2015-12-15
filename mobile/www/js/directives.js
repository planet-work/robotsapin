'use strict';

/* Directives */


angular.module('sapinApp.directives', [])
  .directive('appVersion', ['version',
    function(version) {
      return function(scope, elm, attrs) {
        elm.text(version);
      };
    }
  ])
  .directive('sapinCheck', [
    function() {
      return {
        require: 'ngModel',
        link: function(scope, elem, attrs, ctrl) {
          var me = attrs.ngModel;
          var matchTo = attrs.pwCheck;
                  console.log(me, matchTo)
          scope.$watchGroup([me, matchTo], function(value) {
            ctrl.$setValidity('pwmatch', value[0] === value[1]);
          });

        }
      }
    }
  ]);

