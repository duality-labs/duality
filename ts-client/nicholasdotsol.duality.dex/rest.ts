/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface DexAdjanceyMatrix {
  /** @format uint64 */
  id?: string;
  edgeRow?: DexEdgeRow;
}

export interface DexEdgeRow {
  /** @format uint64 */
  id?: string;
  edge?: boolean;
}

export interface DexFeeList {
  /** @format uint64 */
  id?: string;

  /** @format int64 */
  fee?: string;
}

export interface DexLimitOrderTranche {
  pairId?: string;
  tokenIn?: string;

  /** @format int64 */
  tickIndex?: string;

  /** @format uint64 */
  trancheIndex?: string;
  reservesTokenIn?: string;
  reservesTokenOut?: string;
  totalTokenIn?: string;
  totalTokenOut?: string;
}

export interface DexLimitOrderTrancheTrancheIndexes {
  /** @format uint64 */
  fillTrancheIndex?: string;

  /** @format uint64 */
  placeTrancheIndex?: string;
}

export interface DexLimitOrderTrancheUser {
  pairId?: string;
  token?: string;

  /** @format int64 */
  tickIndex?: string;

  /** @format uint64 */
  count?: string;
  address?: string;
  sharesOwned?: string;
  sharesWithdrawn?: string;
  sharesCancelled?: string;
}

export type DexMsgCancelLimitOrderResponse = object;

export interface DexMsgDepositResponse {
  Reserve0Deposited?: string[];
  Reserve1Deposited?: string[];
}

export type DexMsgPlaceLimitOrderResponse = object;

export interface DexMsgSwapResponse {
  /**
   * Coin defines a token with a denomination and an amount.
   *
   * NOTE: The amount field is an Int which implements the custom method
   * signatures required by gogoproto.
   */
  coinOut?: V1Beta1Coin;
}

export type DexMsgWithdrawFilledLimitOrderResponse = object;

export type DexMsgWithdrawlResponse = object;

export interface DexPairMap {
  pairId?: string;
  tokenPair?: DexTokenPairType;

  /** @format int64 */
  maxTick?: string;

  /** @format int64 */
  minTick?: string;
}

/**
 * Params defines the parameters for the module.
 */
export type DexParams = object;

export interface DexQueryAllAdjanceyMatrixResponse {
  AdjanceyMatrix?: DexAdjanceyMatrix[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface DexQueryAllEdgeRowResponse {
  EdgeRow?: DexEdgeRow[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface DexQueryAllFeeListResponse {
  FeeList?: DexFeeList[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface DexQueryAllLimitOrderTrancheResponse {
  LimitOrderTranche?: DexLimitOrderTranche[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface DexQueryAllLimitOrderTrancheUserResponse {
  LimitOrderTrancheUser?: DexLimitOrderTrancheUser[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface DexQueryAllPairMapResponse {
  pairMap?: DexPairMap[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface DexQueryAllSharesResponse {
  shares?: DexShares[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface DexQueryAllTickMapResponse {
  tickMap?: DexTickMap[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface DexQueryAllTokenMapResponse {
  tokenMap?: DexTokenMap[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface DexQueryAllTokensResponse {
  Tokens?: DexTokens[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface DexQueryGetAdjanceyMatrixResponse {
  AdjanceyMatrix?: DexAdjanceyMatrix;
}

export interface DexQueryGetEdgeRowResponse {
  EdgeRow?: DexEdgeRow;
}

export interface DexQueryGetFeeListResponse {
  FeeList?: DexFeeList;
}

export interface DexQueryGetLimitOrderTrancheResponse {
  LimitOrderTranche?: DexLimitOrderTranche;
}

export interface DexQueryGetLimitOrderTrancheUserResponse {
  LimitOrderTrancheUser?: DexLimitOrderTrancheUser;
}

export interface DexQueryGetPairMapResponse {
  pairMap?: DexPairMap;
}

export interface DexQueryGetSharesResponse {
  shares?: DexShares;
}

export interface DexQueryGetTickMapResponse {
  tickMap?: DexTickMap;
}

export interface DexQueryGetTokenMapResponse {
  tokenMap?: DexTokenMap;
}

export interface DexQueryGetTokensResponse {
  Tokens?: DexTokens;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface DexQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: DexParams;
}

export interface DexReserve0AndSharesType {
  reserve0?: string;
  totalShares?: string;
}

export interface DexShares {
  address?: string;
  pairId?: string;

  /** @format int64 */
  tickIndex?: string;

  /** @format uint64 */
  feeIndex?: string;
  sharesOwned?: string;
}

export interface DexTickDataType {
  reserve0AndShares?: DexReserve0AndSharesType[];
  reserve1?: string[];
}

export interface DexTickMap {
  pairId?: string;

  /** @format int64 */
  tickIndex?: string;
  tickData?: DexTickDataType;
  LimitOrderTranche0to1?: DexLimitOrderTrancheTrancheIndexes;
  LimitOrderTranche1to0?: DexLimitOrderTrancheTrancheIndexes;
}

export interface DexTokenMap {
  address?: string;

  /** @format int64 */
  index?: string;
}

export interface DexTokenPairType {
  /** @format int64 */
  currentTick0To1?: string;

  /** @format int64 */
  currentTick1To0?: string;
}

export interface DexTokens {
  /** @format uint64 */
  id?: string;
  address?: string;
}

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
* Coin defines a token with a denomination and an amount.

NOTE: The amount field is an Int which implements the custom method
signatures required by gogoproto.
*/
export interface V1Beta1Coin {
  denom?: string;
  amount?: string;
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;

  /**
   * reverse is set to true if results are to be returned in the descending order.
   *
   * Since: cosmos-sdk 0.43
   */
  reverse?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /**
   * next_key is the key to be passed to PageRequest.key to
   * query the next page most efficiently
   * @format byte
   */
  next_key?: string;

  /**
   * total is total number of results available if PageRequest.count_total
   * was set, its value is undefined otherwise
   * @format uint64
   */
  total?: string;
}

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.instance.defaults.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title dex/adjancey_matrix.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryAdjanceyMatrixAll
   * @summary Queries a list of AdjanceyMatrix items.
   * @request GET:/NicholasDotSol/duality/dex/adjancey_matrix
   */
  queryAdjanceyMatrixAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DexQueryAllAdjanceyMatrixResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/adjancey_matrix`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryAdjanceyMatrix
   * @summary Queries a AdjanceyMatrix by id.
   * @request GET:/NicholasDotSol/duality/dex/adjancey_matrix/{id}
   */
  queryAdjanceyMatrix = (id: string, params: RequestParams = {}) =>
    this.request<DexQueryGetAdjanceyMatrixResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/adjancey_matrix/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryEdgeRowAll
   * @summary Queries a list of EdgeRow items.
   * @request GET:/NicholasDotSol/duality/dex/edge_row
   */
  queryEdgeRowAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DexQueryAllEdgeRowResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/edge_row`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryEdgeRow
   * @summary Queries a EdgeRow by id.
   * @request GET:/NicholasDotSol/duality/dex/edge_row/{id}
   */
  queryEdgeRow = (id: string, params: RequestParams = {}) =>
    this.request<DexQueryGetEdgeRowResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/edge_row/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryFeeListAll
   * @summary Queries a list of FeeList items.
   * @request GET:/NicholasDotSol/duality/dex/fee_list
   */
  queryFeeListAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DexQueryAllFeeListResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/fee_list`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryFeeList
   * @summary Queries a FeeList by id.
   * @request GET:/NicholasDotSol/duality/dex/fee_list/{id}
   */
  queryFeeList = (id: string, params: RequestParams = {}) =>
    this.request<DexQueryGetFeeListResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/fee_list/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLimitOrderTrancheAll
   * @summary Queries a list of LimitOrderTranche items.
   * @request GET:/NicholasDotSol/duality/dex/limit_order_tranche
   */
  queryLimitOrderTrancheAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DexQueryAllLimitOrderTrancheResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/limit_order_tranche`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLimitOrderTranche
   * @summary Queries a LimitOrderTranche by index.
   * @request GET:/NicholasDotSol/duality/dex/limit_order_tranche/{pairId}/{token}/{tickIndex}/{trancheIndex}
   */
  queryLimitOrderTranche = (
    pairId: string,
    token: string,
    tickIndex: string,
    trancheIndex: string,
    params: RequestParams = {},
  ) =>
    this.request<DexQueryGetLimitOrderTrancheResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/limit_order_tranche/${pairId}/${token}/${tickIndex}/${trancheIndex}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLimitOrderTrancheUserAll
   * @summary Queries a list of LimitOrderTrancheMap items.
   * @request GET:/NicholasDotSol/duality/dex/limit_order_tranche_user
   */
  queryLimitOrderTrancheUserAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DexQueryAllLimitOrderTrancheUserResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/limit_order_tranche_user`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLimitOrderTrancheUser
   * @summary Queries a LimitOrderTrancheUser by index.
   * @request GET:/NicholasDotSol/duality/dex/limit_order_tranche_user/{pairId}/{token}/{tickIndex}/{count}/{address}
   */
  queryLimitOrderTrancheUser = (
    pairId: string,
    token: string,
    tickIndex: string,
    count: string,
    address: string,
    params: RequestParams = {},
  ) =>
    this.request<DexQueryGetLimitOrderTrancheUserResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/limit_order_tranche_user/${pairId}/${token}/${tickIndex}/${count}/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPairMapAll
   * @summary Queries a list of PairMap items.
   * @request GET:/NicholasDotSol/duality/dex/pair_map
   */
  queryPairMapAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DexQueryAllPairMapResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/pair_map`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPairMap
   * @summary Queries a PairMap by index.
   * @request GET:/NicholasDotSol/duality/dex/pair_map/{pairId}
   */
  queryPairMap = (pairId: string, params: RequestParams = {}) =>
    this.request<DexQueryGetPairMapResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/pair_map/${pairId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/NicholasDotSol/duality/dex/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<DexQueryParamsResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QuerySharesAll
   * @summary Queries a list of Shares items.
   * @request GET:/NicholasDotSol/duality/dex/shares
   */
  querySharesAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DexQueryAllSharesResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/shares`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryShares
   * @summary Queries a Shares by index.
   * @request GET:/NicholasDotSol/duality/dex/shares/{address}/{pairId}/{tickIndex}/{fee}
   */
  queryShares = (address: string, pairId: string, tickIndex: string, fee: string, params: RequestParams = {}) =>
    this.request<DexQueryGetSharesResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/shares/${address}/${pairId}/${tickIndex}/${fee}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTickMapAll
   * @summary Queries a list of TickMap items.
   * @request GET:/NicholasDotSol/duality/dex/tick_map
   */
  queryTickMapAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DexQueryAllTickMapResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/tick_map`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTickMap
   * @summary Queries a TickMap by index.
   * @request GET:/NicholasDotSol/duality/dex/tick_map/{pairId}/{tickIndex}
   */
  queryTickMap = (pairId: string, tickIndex: string, params: RequestParams = {}) =>
    this.request<DexQueryGetTickMapResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/tick_map/${pairId}/${tickIndex}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTokenMapAll
   * @summary Queries a list of TokenMap items.
   * @request GET:/NicholasDotSol/duality/dex/token_map
   */
  queryTokenMapAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DexQueryAllTokenMapResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/token_map`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTokenMap
   * @summary Queries a TokenMap by index.
   * @request GET:/NicholasDotSol/duality/dex/token_map/{address}
   */
  queryTokenMap = (address: string, params: RequestParams = {}) =>
    this.request<DexQueryGetTokenMapResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/token_map/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTokensAll
   * @summary Queries a list of Tokens items.
   * @request GET:/NicholasDotSol/duality/dex/tokens
   */
  queryTokensAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DexQueryAllTokensResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/tokens`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryTokens
   * @summary Queries a Tokens by id.
   * @request GET:/NicholasDotSol/duality/dex/tokens/{id}
   */
  queryTokens = (id: string, params: RequestParams = {}) =>
    this.request<DexQueryGetTokensResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/tokens/${id}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
