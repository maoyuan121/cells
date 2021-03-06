/**
 * Pydio Cells Rest API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */

'use strict';

exports.__esModule = true;

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

var _ApiClient = require('../ApiClient');

var _ApiClient2 = _interopRequireDefault(_ApiClient);

/**
* The RestSettingsEntryMeta model module.
* @module model/RestSettingsEntryMeta
* @version 1.0
*/

var RestSettingsEntryMeta = (function () {
    /**
    * Constructs a new <code>RestSettingsEntryMeta</code>.
    * @alias module:model/RestSettingsEntryMeta
    * @class
    */

    function RestSettingsEntryMeta() {
        _classCallCheck(this, RestSettingsEntryMeta);

        this.IconClass = undefined;
        this.Component = undefined;
        this.Props = undefined;
        this.Advanced = undefined;
        this.Indexed = undefined;
    }

    /**
    * Constructs a <code>RestSettingsEntryMeta</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/RestSettingsEntryMeta} obj Optional instance to populate.
    * @return {module:model/RestSettingsEntryMeta} The populated <code>RestSettingsEntryMeta</code> instance.
    */

    RestSettingsEntryMeta.constructFromObject = function constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RestSettingsEntryMeta();

            if (data.hasOwnProperty('IconClass')) {
                obj['IconClass'] = _ApiClient2['default'].convertToType(data['IconClass'], 'String');
            }
            if (data.hasOwnProperty('Component')) {
                obj['Component'] = _ApiClient2['default'].convertToType(data['Component'], 'String');
            }
            if (data.hasOwnProperty('Props')) {
                obj['Props'] = _ApiClient2['default'].convertToType(data['Props'], 'String');
            }
            if (data.hasOwnProperty('Advanced')) {
                obj['Advanced'] = _ApiClient2['default'].convertToType(data['Advanced'], 'Boolean');
            }
            if (data.hasOwnProperty('Indexed')) {
                obj['Indexed'] = _ApiClient2['default'].convertToType(data['Indexed'], ['String']);
            }
        }
        return obj;
    };

    /**
    * @member {String} IconClass
    */
    return RestSettingsEntryMeta;
})();

exports['default'] = RestSettingsEntryMeta;
module.exports = exports['default'];

/**
* @member {String} Component
*/

/**
* @member {String} Props
*/

/**
* @member {Boolean} Advanced
*/

/**
* @member {Array.<String>} Indexed
*/
