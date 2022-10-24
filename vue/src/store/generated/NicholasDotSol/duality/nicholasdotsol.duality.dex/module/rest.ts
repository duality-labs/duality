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

export interface DexLimitOrderPool {
  /** @format uint64 */
  count?: string;

  /** @format uint64 */
  currentLimitOrderKey?: string;
}

export interface DexLimitOrderPoolFillMap {
  pairId?: string;
  token?: string;

  /** @format int64 */
  tickIndex?: string;

  /** @format uint64 */
  count?: string;
  fill?: string;
}

export interface DexLimitOrderPoolReserveMap {
  pairId?: string;
  token?: string;

  /** @format int64 */
  tickIndex?: string;

  /** @format uint64 */
  count?: string;
  reserves?: string;
}

export interface DexLimitOrderPoolTotalSharesMap {
  pairId?: string;
  token?: string;

  /** @format int64 */
  tickIndex?: string;

  /** @format uint64 */
  count?: string;
  totalShares?: string;
}

export interface DexLimitOrderPoolUserShareMap {
  pairId?: string;
  token?: string;

  /** @format int64 */
  tickIndex?: string;

  /** @format uint64 */
  count?: string;
  address?: string;
  sharesOwned?: string;
}

export interface DexLimitOrderPoolUserSharesWithdrawn {
  pairId?: string;
  token?: string;

  /** @format int64 */
  tickIndex?: string;

  /** @format uint64 */
  count?: string;
  address?: string;
  sharesWithdrawn?: string;
}

export type DexMsgCancelLimitOrderResponse = object;

export type DexMsgDepositResponse = object;

export type DexMsgPlaceLimitOrderResponse = object;

export type DexMsgSwapResponse = object;

export type DexMsgWithdrawFilledLimitOrderResponse = object;

export type DexMsgWithdrawlResponse = object;

export interface DexPairMap {
  pairId?: string;
  tokenPair?: DexTokenPairType;

  /** @format int64 */
  totalTickCount?: string;
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

export interface DexQueryAllLimitOrderPoolFillMapResponse {
  limitOrderPoolFillMap?: DexLimitOrderPoolFillMap[];

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

export interface DexQueryAllLimitOrderPoolReserveMapResponse {
  limitOrderPoolReserveMap?: DexLimitOrderPoolReserveMap[];

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

export interface DexQueryAllLimitOrderPoolTotalSharesMapResponse {
  limitOrderPoolTotalSharesMap?: DexLimitOrderPoolTotalSharesMap[];

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

export interface DexQueryAllLimitOrderPoolUserShareMapResponse {
  limitOrderPoolUserShareMap?: DexLimitOrderPoolUserShareMap[];

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

export interface DexQueryAllLimitOrderPoolUserSharesWithdrawnResponse {
  limitOrderPoolUserSharesWithdrawn?: DexLimitOrderPoolUserSharesWithdrawn[];

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

export interface DexQueryGetLimitOrderPoolFillMapResponse {
  limitOrderPoolFillMap?: DexLimitOrderPoolFillMap;
}

export interface DexQueryGetLimitOrderPoolReserveMapResponse {
  limitOrderPoolReserveMap?: DexLimitOrderPoolReserveMap;
}

export interface DexQueryGetLimitOrderPoolTotalSharesMapResponse {
  limitOrderPoolTotalSharesMap?: DexLimitOrderPoolTotalSharesMap;
}

export interface DexQueryGetLimitOrderPoolUserShareMapResponse {
  limitOrderPoolUserShareMap?: DexLimitOrderPoolUserShareMap;
}

export interface DexQueryGetLimitOrderPoolUserSharesWithdrawnResponse {
  limitOrderPoolUserSharesWithdrawn?: DexLimitOrderPoolUserSharesWithdrawn;
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
  /** Params defines the parameters for the module. */
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
  LimitOrderPool0to1?: DexLimitOrderPool;
  LimitOrderPool1to0?: DexLimitOrderPool;
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
  /** @format byte */
  next_key?: string;

  /** @format uint64 */
  total?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: keyof Omit<Body, "body" | "bodyUsed">;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
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
   * @name QueryLimitOrderPoolFillMapAll
   * @summary Queries a list of LimitOrderPoolFillMap items.
   * @request GET:/NicholasDotSol/duality/dex/limit_order_pool_fill_map
   */
  queryLimitOrderPoolFillMapAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DexQueryAllLimitOrderPoolFillMapResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/limit_order_pool_fill_map`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLimitOrderPoolFillMap
   * @summary Queries a LimitOrderPoolFillMap by index.
   * @request GET:/NicholasDotSol/duality/dex/limit_order_pool_fill_map/{pairId}/{token}/{tickIndex}/{count}
   */
  queryLimitOrderPoolFillMap = (
    pairId: string,
    token: string,
    tickIndex: string,
    count: string,
    params: RequestParams = {},
  ) =>
    this.request<DexQueryGetLimitOrderPoolFillMapResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/limit_order_pool_fill_map/${pairId}/${token}/${tickIndex}/${count}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLimitOrderPoolReserveMapAll
   * @summary Queries a list of LimitOrderPoolReserveMap items.
   * @request GET:/NicholasDotSol/duality/dex/limit_order_pool_reserve_map
   */
  queryLimitOrderPoolReserveMapAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DexQueryAllLimitOrderPoolReserveMapResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/limit_order_pool_reserve_map`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLimitOrderPoolReserveMap
   * @summary Queries a LimitOrderPoolReserveMap by index.
   * @request GET:/NicholasDotSol/duality/dex/limit_order_pool_reserve_map/{pairId}/{token}/{tickIndex}/{count}
   */
  queryLimitOrderPoolReserveMap = (
    pairId: string,
    token: string,
    tickIndex: string,
    count: string,
    params: RequestParams = {},
  ) =>
    this.request<DexQueryGetLimitOrderPoolReserveMapResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/limit_order_pool_reserve_map/${pairId}/${token}/${tickIndex}/${count}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLimitOrderPoolTotalSharesMapAll
   * @summary Queries a list of LimitOrderPoolTotalSharesMap items.
   * @request GET:/NicholasDotSol/duality/dex/limit_order_pool_total_shares_map
   */
  queryLimitOrderPoolTotalSharesMapAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DexQueryAllLimitOrderPoolTotalSharesMapResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/limit_order_pool_total_shares_map`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLimitOrderPoolTotalSharesMap
   * @summary Queries a LimitOrderPoolTotalSharesMap by index.
   * @request GET:/NicholasDotSol/duality/dex/limit_order_pool_total_shares_map/{pairId}/{token}/{tickIndex}/{count}
   */
  queryLimitOrderPoolTotalSharesMap = (
    pairId: string,
    token: string,
    tickIndex: string,
    count: string,
    params: RequestParams = {},
  ) =>
    this.request<DexQueryGetLimitOrderPoolTotalSharesMapResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/limit_order_pool_total_shares_map/${pairId}/${token}/${tickIndex}/${count}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLimitOrderPoolUserShareMapAll
   * @summary Queries a list of LimitOrderPoolUserShareMap items.
   * @request GET:/NicholasDotSol/duality/dex/limit_order_pool_user_share_map
   */
  queryLimitOrderPoolUserShareMapAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DexQueryAllLimitOrderPoolUserShareMapResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/limit_order_pool_user_share_map`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLimitOrderPoolUserShareMap
   * @summary Queries a LimitOrderPoolUserShareMap by index.
   * @request GET:/NicholasDotSol/duality/dex/limit_order_pool_user_share_map/{pairId}/{token}/{tickIndex}/{count}/{address}
   */
  queryLimitOrderPoolUserShareMap = (
    pairId: string,
    token: string,
    tickIndex: string,
    count: string,
    address: string,
    params: RequestParams = {},
  ) =>
    this.request<DexQueryGetLimitOrderPoolUserShareMapResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/limit_order_pool_user_share_map/${pairId}/${token}/${tickIndex}/${count}/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLimitOrderPoolUserSharesWithdrawnAll
   * @summary Queries a list of LimitOrderPoolUserSharesWithdrawn items.
   * @request GET:/NicholasDotSol/duality/dex/limit_order_pool_user_shares_withdrawn
   */
  queryLimitOrderPoolUserSharesWithdrawnAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DexQueryAllLimitOrderPoolUserSharesWithdrawnResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/limit_order_pool_user_shares_withdrawn`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLimitOrderPoolUserSharesWithdrawn
   * @summary Queries a LimitOrderPoolUserSharesWithdrawn by index.
   * @request GET:/NicholasDotSol/duality/dex/limit_order_pool_user_shares_withdrawn/{pairId}/{token}/{tickIndex}/{count}/{address}
   */
  queryLimitOrderPoolUserSharesWithdrawn = (
    pairId: string,
    token: string,
    tickIndex: string,
    count: string,
    address: string,
    params: RequestParams = {},
  ) =>
    this.request<DexQueryGetLimitOrderPoolUserSharesWithdrawnResponse, RpcStatus>({
      path: `/NicholasDotSol/duality/dex/limit_order_pool_user_shares_withdrawn/${pairId}/${token}/${tickIndex}/${count}/${address}`,
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
