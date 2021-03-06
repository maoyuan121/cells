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
* The RestGetBulkMetaRequest model module.
* @module model/RestGetBulkMetaRequest
* @version 1.0
*/

var RestGetBulkMetaRequest = (function () {
    /**
    * Constructs a new <code>RestGetBulkMetaRequest</code>.
    * @alias module:model/RestGetBulkMetaRequest
    * @class
    */

    function RestGetBulkMetaRequest() {
        _classCallCheck(this, RestGetBulkMetaRequest);

        this.NodePaths = undefined;
        this.AllMetaProviders = undefined;
        this.Versions = undefined;
        this.Offset = undefined;
        this.Limit = undefined;
    }

    /**
    * Constructs a <code>RestGetBulkMetaRequest</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/RestGetBulkMetaRequest} obj Optional instance to populate.
    * @return {module:model/RestGetBulkMetaRequest} The populated <code>RestGetBulkMetaRequest</code> instance.
    */

    RestGetBulkMetaRequest.constructFromObject = function constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RestGetBulkMetaRequest();

            if (data.hasOwnProperty('NodePaths')) {
                obj['NodePaths'] = _ApiClient2['default'].convertToType(data['NodePaths'], ['String']);
            }
            if (data.hasOwnProperty('AllMetaProviders')) {
                obj['AllMetaProviders'] = _ApiClient2['default'].convertToType(data['AllMetaProviders'], 'Boolean');
            }
            if (data.hasOwnProperty('Versions')) {
                obj['Versions'] = _ApiClient2['default'].convertToType(data['Versions'], 'Boolean');
            }
            if (data.hasOwnProperty('Offset')) {
                obj['Offset'] = _ApiClient2['default'].convertToType(data['Offset'], 'Number');
            }
            if (data.hasOwnProperty('Limit')) {
                obj['Limit'] = _ApiClient2['default'].convertToType(data['Limit'], 'Number');
            }
        }
        return obj;
    };

    /**
    * @member {Array.<String>} NodePaths
    */
    return RestGetBulkMetaRequest;
})();

exports['default'] = RestGetBulkMetaRequest;
module.exports = exports['default'];

/**
* @member {Boolean} AllMetaProviders
*/

/**
* @member {Boolean} Versions
*/

/**
* @member {Number} Offset
*/

/**
* @member {Number} Limit
*/
