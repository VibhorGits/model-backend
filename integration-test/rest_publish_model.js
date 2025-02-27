import http from "k6/http";
import {
  check,
  group,
  sleep
} from "k6";
import {
  FormData
} from "https://jslib.k6.io/formdata/0.0.2/index.js";
import {
  randomString
} from "https://jslib.k6.io/k6-utils/1.1.0/index.js";

import {
  genHeader,
  isValidOwner,
} from "./helpers.js";

import * as constant from "./const.js"

const model_def_name = "model-definitions/local"

export function PublishUnpublishModel(header) {
  // Model Backend API: PublishModel
  {
    group("Model Backend API: PublishModel", function () {
      let fd_cls = new FormData();
      let model_id = randomString(10)
      let model_description = randomString(20)
      fd_cls.append("id", model_id);
      fd_cls.append("description", model_description);
      fd_cls.append("model_definition", model_def_name);
      fd_cls.append("content", http.file(constant.cls_model, "dummy-cls-model.zip"));
      let createClsModelRes = http.request("POST", `${constant.apiPublicHost}/v1alpha/${constant.namespace}/models/multipart`, fd_cls.body(), {
        headers: genHeader(`multipart/form-data; boundary=${fd_cls.boundary}`, header.headers.Authorization),
      })
      check(createClsModelRes, {
        "POST /v1alpha/models/multipart task cls response status": (r) =>
          r.status === 201,
        "POST /v1alpha/models/multipart task cls response operation.name": (r) =>
          r.json().operation.name !== undefined,
      });

      // Check model creation finished
      let currentTime = new Date().getTime();
      let timeoutTime = new Date().getTime() + 120000;
      while (timeoutTime > currentTime) {
        let res = http.get(`${constant.apiPublicHost}/v1alpha/${createClsModelRes.json().operation.name}`, header)
        if (res.json().operation.done === true) {
          break
        }
        sleep(1)
        currentTime = new Date().getTime();
      }

      check(http.post(`${constant.apiPublicHost}/v1alpha/${constant.namespace}/models/${model_id}/publish`, null, header), {
        [`POST /v1alpha/models/${model_id}/publish task cls response status`]: (r) =>
          r.status === 200,
        [`POST /v1alpha/models/${model_id}/publish task cls response model.name`]: (r) =>
          r.json().model.name === `${constant.namespace}/models/${model_id}`,
        [`POST /v1alpha/models/${model_id}/publish task cls response model.uid`]: (r) =>
          r.json().model.uid !== undefined,
        [`POST /v1alpha/models/${model_id}/publish task cls response model.id`]: (r) =>
          r.json().model.id === model_id,
        [`POST /v1alpha/models/${model_id}/publish task cls response model.description`]: (r) =>
          r.json().model.description === model_description,
        [`POST /v1alpha/models/${model_id}/publish task cls response model.model_definition`]: (r) =>
          r.json().model.model_definition === model_def_name,
        [`POST /v1alpha/models/${model_id}/publish task cls response model.configuration`]: (r) =>
          r.json().model.configuration !== undefined,
        [`POST /v1alpha/models/${model_id}/publish task cls response model.visibility`]: (r) =>
          r.json().model.visibility === "VISIBILITY_PUBLIC",
        [`POST /v1alpha/models/${model_id}/publish task cls response model.owner`]: (r) =>
          isValidOwner(r.json().model.user),
        [`POST /v1alpha/models/${model_id}/publish task cls response model.create_time`]: (r) =>
          r.json().model.create_time !== undefined,
        [`POST /v1alpha/models/${model_id}/publish task cls response model.update_time`]: (r) =>
          r.json().model.update_time !== undefined,
      });

      check(http.post(`${constant.apiPublicHost}/v1alpha/${constant.namespace}/models/${model_id}/unpublish`, null, header), {
        [`POST /v1alpha/models/${model_id}/unpublish task cls response status`]: (r) =>
          r.status === 200,
        [`POST /v1alpha/models/${model_id}/unpublish task cls response model.name`]: (r) =>
          r.json().model.name === `${constant.namespace}/models/${model_id}`,
        [`POST /v1alpha/models/${model_id}/unpublish task cls response model.uid`]: (r) =>
          r.json().model.uid !== undefined,
        [`POST /v1alpha/models/${model_id}/unpublish task cls response model.id`]: (r) =>
          r.json().model.id === model_id,
        [`POST /v1alpha/models/${model_id}/unpublish task cls response model.description`]: (r) =>
          r.json().model.description === model_description,
        [`POST /v1alpha/models/${model_id}/unpublish task cls response model.model_definition`]: (r) =>
          r.json().model.model_definition === model_def_name,
        [`POST /v1alpha/models/${model_id}/unpublish task cls response model.configuration`]: (r) =>
          r.json().model.configuration !== undefined,
        [`POST /v1alpha/models/${model_id}/unpublish task cls response model.visibility`]: (r) =>
          r.json().model.visibility === "VISIBILITY_PRIVATE",
        [`POST /v1alpha/models/${model_id}/unpublish task cls response model.owner`]: (r) =>
          isValidOwner(r.json().model.user),
        [`POST /v1alpha/models/${model_id}/unpublish task cls response model.create_time`]: (r) =>
          r.json().model.create_time !== undefined,
        [`POST /v1alpha/models/${model_id}/unpublish task cls response model.update_time`]: (r) =>
          r.json().model.update_time !== undefined,
      });

      check(http.post(`${constant.apiPublicHost}/v1alpha/${constant.namespace}/models/${randomString(10)}/publish`, null, header), {
        [`POST /v1alpha/models/${model_id}/publish task cls response not found status`]: (r) => r.status === 404,
      });

      check(http.post(`${constant.apiPublicHost}/v1alpha/${constant.namespace}/models/${randomString(10)}/unpublish`, null, header), {
        [`POST /v1alpha/models/${model_id}/unpublish task cls response not found status`]: (r) => r.status === 404,
      });

      currentTime = new Date().getTime();
      timeoutTime = new Date().getTime() + 120000;
      while (timeoutTime > currentTime) {
        let res = http.get(`${constant.apiPublicHost}/v1alpha/${constant.namespace}/models/${model_id}/watch`, header)
        if (res.json().state !== "STATE_UNSPECIFIED") {
          break
        }
        sleep(1)
        currentTime = new Date().getTime();
      }

      // clean up
      check(http.request("DELETE", `${constant.apiPublicHost}/v1alpha/${constant.namespace}/models/${model_id}`, null, header), {
        "DELETE clean up response status": (r) =>
          r.status === 204
      });
    });
  }

}
