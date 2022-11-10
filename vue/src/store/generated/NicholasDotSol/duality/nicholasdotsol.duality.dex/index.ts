import { txClient, queryClient, MissingWalletError , registry} from './module'

import { FeeList } from "./module/types/dex/fee_list"
import { LimitOrderPool } from "./module/types/dex/limit_order_pool"
import { LimitOrderPoolFillObject } from "./module/types/dex/limit_order_pool_fill_object"
import { LimitOrderPoolReserveObject } from "./module/types/dex/limit_order_pool_reserve_object"
import { LimitOrderPoolTotalSharesObject } from "./module/types/dex/limit_order_pool_total_shares_object"
import { LimitOrderPoolUserShareObject } from "./module/types/dex/limit_order_pool_user_share_object"
import { LimitOrderPoolUserSharesWithdrawnObject } from "./module/types/dex/limit_order_pool_user_shares_withdrawn_object"
import { PairObject } from "./module/types/dex/pair_object"
import { Params } from "./module/types/dex/params"
import { Reserve0AndSharesType } from "./module/types/dex/reserve_0_and_shares_type"
import { Shares } from "./module/types/dex/shares"
import { TickDataType } from "./module/types/dex/tick_data_type"
import { TickObject } from "./module/types/dex/tick_object"
import { TokenObject } from "./module/types/dex/token_object"
import { TokenPairType } from "./module/types/dex/token_pair_type"
import { Tokens } from "./module/types/dex/tokens"


export { FeeList, LimitOrderPool, LimitOrderPoolFillObject, LimitOrderPoolReserveObject, LimitOrderPoolTotalSharesObject, LimitOrderPoolUserShareObject, LimitOrderPoolUserSharesWithdrawnObject, PairObject, Params, Reserve0AndSharesType, Shares, TickDataType, TickObject, TokenObject, TokenPairType, Tokens };

async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
		structure.fields.push(field)
	}
	return structure
}

const getDefaultState = () => {
	return {
				Params: {},
				TickObject: {},
				TickObjectAll: {},
				PairObject: {},
				PairObjectAll: {},
				Tokens: {},
				TokensAll: {},
				TokenObject: {},
				TokenObjectAll: {},
				Shares: {},
				SharesAll: {},
				FeeList: {},
				FeeListAll: {},
				LimitOrderPoolUserShareObject: {},
				LimitOrderPoolUserShareObjectAll: {},
				LimitOrderPoolUserSharesWithdrawnObject: {},
				LimitOrderPoolUserSharesWithdrawnObjectAll: {},
				LimitOrderPoolTotalSharesObject: {},
				LimitOrderPoolTotalSharesObjectAll: {},
				LimitOrderPoolReserveObject: {},
				LimitOrderPoolReserveObjectAll: {},
				LimitOrderPoolFillObject: {},
				LimitOrderPoolFillObjectAll: {},
				
				_Structure: {
						FeeList: getStructure(FeeList.fromPartial({})),
						LimitOrderPool: getStructure(LimitOrderPool.fromPartial({})),
						LimitOrderPoolFillObject: getStructure(LimitOrderPoolFillObject.fromPartial({})),
						LimitOrderPoolReserveObject: getStructure(LimitOrderPoolReserveObject.fromPartial({})),
						LimitOrderPoolTotalSharesObject: getStructure(LimitOrderPoolTotalSharesObject.fromPartial({})),
						LimitOrderPoolUserShareObject: getStructure(LimitOrderPoolUserShareObject.fromPartial({})),
						LimitOrderPoolUserSharesWithdrawnObject: getStructure(LimitOrderPoolUserSharesWithdrawnObject.fromPartial({})),
						PairObject: getStructure(PairObject.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						Reserve0AndSharesType: getStructure(Reserve0AndSharesType.fromPartial({})),
						Shares: getStructure(Shares.fromPartial({})),
						TickDataType: getStructure(TickDataType.fromPartial({})),
						TickObject: getStructure(TickObject.fromPartial({})),
						TokenObject: getStructure(TokenObject.fromPartial({})),
						TokenPairType: getStructure(TokenPairType.fromPartial({})),
						Tokens: getStructure(Tokens.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getTickObject: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TickObject[JSON.stringify(params)] ?? {}
		},
				getTickObjectAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TickObjectAll[JSON.stringify(params)] ?? {}
		},
				getPairObject: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PairObject[JSON.stringify(params)] ?? {}
		},
				getPairObjectAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PairObjectAll[JSON.stringify(params)] ?? {}
		},
				getTokens: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Tokens[JSON.stringify(params)] ?? {}
		},
				getTokensAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TokensAll[JSON.stringify(params)] ?? {}
		},
				getTokenObject: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TokenObject[JSON.stringify(params)] ?? {}
		},
				getTokenObjectAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TokenObjectAll[JSON.stringify(params)] ?? {}
		},
				getShares: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Shares[JSON.stringify(params)] ?? {}
		},
				getSharesAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SharesAll[JSON.stringify(params)] ?? {}
		},
				getFeeList: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.FeeList[JSON.stringify(params)] ?? {}
		},
				getFeeListAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.FeeListAll[JSON.stringify(params)] ?? {}
		},
				getLimitOrderPoolUserShareObject: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LimitOrderPoolUserShareObject[JSON.stringify(params)] ?? {}
		},
				getLimitOrderPoolUserShareObjectAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LimitOrderPoolUserShareObjectAll[JSON.stringify(params)] ?? {}
		},
				getLimitOrderPoolUserSharesWithdrawnObject: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LimitOrderPoolUserSharesWithdrawnObject[JSON.stringify(params)] ?? {}
		},
				getLimitOrderPoolUserSharesWithdrawnObjectAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LimitOrderPoolUserSharesWithdrawnObjectAll[JSON.stringify(params)] ?? {}
		},
				getLimitOrderPoolTotalSharesObject: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LimitOrderPoolTotalSharesObject[JSON.stringify(params)] ?? {}
		},
				getLimitOrderPoolTotalSharesObjectAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LimitOrderPoolTotalSharesObjectAll[JSON.stringify(params)] ?? {}
		},
				getLimitOrderPoolReserveObject: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LimitOrderPoolReserveObject[JSON.stringify(params)] ?? {}
		},
				getLimitOrderPoolReserveObjectAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LimitOrderPoolReserveObjectAll[JSON.stringify(params)] ?? {}
		},
				getLimitOrderPoolFillObject: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LimitOrderPoolFillObject[JSON.stringify(params)] ?? {}
		},
				getLimitOrderPoolFillObjectAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LimitOrderPoolFillObjectAll[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: nicholasdotsol.duality.dex initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTickObject({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryTickObject( key.pairId,  key.tickIndex)).data
				
					
				commit('QUERY', { query: 'TickObject', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTickObject', payload: { options: { all }, params: {...key},query }})
				return getters['getTickObject']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTickObject API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTickObjectAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryTickObjectAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryTickObjectAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'TickObjectAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTickObjectAll', payload: { options: { all }, params: {...key},query }})
				return getters['getTickObjectAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTickObjectAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPairObject({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryPairObject( key.pairId)).data
				
					
				commit('QUERY', { query: 'PairObject', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPairObject', payload: { options: { all }, params: {...key},query }})
				return getters['getPairObject']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPairObject API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPairObjectAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryPairObjectAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryPairObjectAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'PairObjectAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPairObjectAll', payload: { options: { all }, params: {...key},query }})
				return getters['getPairObjectAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPairObjectAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTokens({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryTokens( key.id)).data
				
					
				commit('QUERY', { query: 'Tokens', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTokens', payload: { options: { all }, params: {...key},query }})
				return getters['getTokens']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTokens API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTokensAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryTokensAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryTokensAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'TokensAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTokensAll', payload: { options: { all }, params: {...key},query }})
				return getters['getTokensAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTokensAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTokenObject({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryTokenObject( key.address)).data
				
					
				commit('QUERY', { query: 'TokenObject', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTokenObject', payload: { options: { all }, params: {...key},query }})
				return getters['getTokenObject']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTokenObject API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTokenObjectAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryTokenObjectAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryTokenObjectAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'TokenObjectAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTokenObjectAll', payload: { options: { all }, params: {...key},query }})
				return getters['getTokenObjectAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTokenObjectAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryShares({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryShares( key.address,  key.pairId,  key.tickIndex,  key.fee)).data
				
					
				commit('QUERY', { query: 'Shares', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryShares', payload: { options: { all }, params: {...key},query }})
				return getters['getShares']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryShares API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QuerySharesAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.querySharesAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.querySharesAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'SharesAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySharesAll', payload: { options: { all }, params: {...key},query }})
				return getters['getSharesAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QuerySharesAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryFeeList({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryFeeList( key.id)).data
				
					
				commit('QUERY', { query: 'FeeList', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryFeeList', payload: { options: { all }, params: {...key},query }})
				return getters['getFeeList']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryFeeList API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryFeeListAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryFeeListAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryFeeListAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'FeeListAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryFeeListAll', payload: { options: { all }, params: {...key},query }})
				return getters['getFeeListAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryFeeListAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLimitOrderPoolUserShareObject({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLimitOrderPoolUserShareObject( key.pairId,  key.token,  key.tickIndex,  key.count,  key.address)).data
				
					
				commit('QUERY', { query: 'LimitOrderPoolUserShareObject', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLimitOrderPoolUserShareObject', payload: { options: { all }, params: {...key},query }})
				return getters['getLimitOrderPoolUserShareObject']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLimitOrderPoolUserShareObject API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLimitOrderPoolUserShareObjectAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLimitOrderPoolUserShareObjectAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryLimitOrderPoolUserShareObjectAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'LimitOrderPoolUserShareObjectAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLimitOrderPoolUserShareObjectAll', payload: { options: { all }, params: {...key},query }})
				return getters['getLimitOrderPoolUserShareObjectAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLimitOrderPoolUserShareObjectAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLimitOrderPoolUserSharesWithdrawnObject({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLimitOrderPoolUserSharesWithdrawnObject( key.pairId,  key.token,  key.tickIndex,  key.count,  key.address)).data
				
					
				commit('QUERY', { query: 'LimitOrderPoolUserSharesWithdrawnObject', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLimitOrderPoolUserSharesWithdrawnObject', payload: { options: { all }, params: {...key},query }})
				return getters['getLimitOrderPoolUserSharesWithdrawnObject']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLimitOrderPoolUserSharesWithdrawnObject API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLimitOrderPoolUserSharesWithdrawnObjectAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLimitOrderPoolUserSharesWithdrawnObjectAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryLimitOrderPoolUserSharesWithdrawnObjectAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'LimitOrderPoolUserSharesWithdrawnObjectAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLimitOrderPoolUserSharesWithdrawnObjectAll', payload: { options: { all }, params: {...key},query }})
				return getters['getLimitOrderPoolUserSharesWithdrawnObjectAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLimitOrderPoolUserSharesWithdrawnObjectAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLimitOrderPoolTotalSharesObject({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLimitOrderPoolTotalSharesObject( key.pairId,  key.token,  key.tickIndex,  key.count)).data
				
					
				commit('QUERY', { query: 'LimitOrderPoolTotalSharesObject', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLimitOrderPoolTotalSharesObject', payload: { options: { all }, params: {...key},query }})
				return getters['getLimitOrderPoolTotalSharesObject']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLimitOrderPoolTotalSharesObject API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLimitOrderPoolTotalSharesObjectAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLimitOrderPoolTotalSharesObjectAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryLimitOrderPoolTotalSharesObjectAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'LimitOrderPoolTotalSharesObjectAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLimitOrderPoolTotalSharesObjectAll', payload: { options: { all }, params: {...key},query }})
				return getters['getLimitOrderPoolTotalSharesObjectAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLimitOrderPoolTotalSharesObjectAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLimitOrderPoolReserveObject({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLimitOrderPoolReserveObject( key.pairId,  key.token,  key.tickIndex,  key.count)).data
				
					
				commit('QUERY', { query: 'LimitOrderPoolReserveObject', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLimitOrderPoolReserveObject', payload: { options: { all }, params: {...key},query }})
				return getters['getLimitOrderPoolReserveObject']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLimitOrderPoolReserveObject API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLimitOrderPoolReserveObjectAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLimitOrderPoolReserveObjectAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryLimitOrderPoolReserveObjectAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'LimitOrderPoolReserveObjectAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLimitOrderPoolReserveObjectAll', payload: { options: { all }, params: {...key},query }})
				return getters['getLimitOrderPoolReserveObjectAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLimitOrderPoolReserveObjectAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLimitOrderPoolFillObject({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLimitOrderPoolFillObject( key.pairId,  key.token,  key.tickIndex,  key.count)).data
				
					
				commit('QUERY', { query: 'LimitOrderPoolFillObject', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLimitOrderPoolFillObject', payload: { options: { all }, params: {...key},query }})
				return getters['getLimitOrderPoolFillObject']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLimitOrderPoolFillObject API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLimitOrderPoolFillObjectAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryLimitOrderPoolFillObjectAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryLimitOrderPoolFillObjectAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'LimitOrderPoolFillObjectAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLimitOrderPoolFillObjectAll', payload: { options: { all }, params: {...key},query }})
				return getters['getLimitOrderPoolFillObjectAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLimitOrderPoolFillObjectAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgDeposit({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeposit(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeposit:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeposit:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgWithdrawFilledLimitOrder({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgWithdrawFilledLimitOrder(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdrawFilledLimitOrder:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgWithdrawFilledLimitOrder:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgPlaceLimitOrder({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgPlaceLimitOrder(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPlaceLimitOrder:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgPlaceLimitOrder:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgWithdrawl({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgWithdrawl(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdrawl:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgWithdrawl:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCancelLimitOrder({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCancelLimitOrder(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancelLimitOrder:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCancelLimitOrder:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSwap({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSwap(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSwap:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSwap:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgDeposit({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeposit(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeposit:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeposit:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgWithdrawFilledLimitOrder({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgWithdrawFilledLimitOrder(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdrawFilledLimitOrder:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgWithdrawFilledLimitOrder:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgPlaceLimitOrder({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgPlaceLimitOrder(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPlaceLimitOrder:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgPlaceLimitOrder:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgWithdrawl({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgWithdrawl(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdrawl:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgWithdrawl:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCancelLimitOrder({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCancelLimitOrder(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancelLimitOrder:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCancelLimitOrder:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSwap({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgSwap(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSwap:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSwap:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
