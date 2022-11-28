import { Client, registry, MissingWalletError } from 'NicholasDotSol-duality-client-ts'

import { AdjanceyMatrix } from "NicholasDotSol-duality-client-ts/nicholasdotsol.duality.dex/types"
import { EdgeRow } from "NicholasDotSol-duality-client-ts/nicholasdotsol.duality.dex/types"
import { FeeList } from "NicholasDotSol-duality-client-ts/nicholasdotsol.duality.dex/types"
import { LimitOrderTrancheTrancheIndexes } from "NicholasDotSol-duality-client-ts/nicholasdotsol.duality.dex/types"
import { LimitOrderTranche } from "NicholasDotSol-duality-client-ts/nicholasdotsol.duality.dex/types"
import { LimitOrderTrancheUser } from "NicholasDotSol-duality-client-ts/nicholasdotsol.duality.dex/types"
import { PairMap } from "NicholasDotSol-duality-client-ts/nicholasdotsol.duality.dex/types"
import { Params } from "NicholasDotSol-duality-client-ts/nicholasdotsol.duality.dex/types"
import { Reserve0AndSharesType } from "NicholasDotSol-duality-client-ts/nicholasdotsol.duality.dex/types"
import { Shares } from "NicholasDotSol-duality-client-ts/nicholasdotsol.duality.dex/types"
import { TickDataType } from "NicholasDotSol-duality-client-ts/nicholasdotsol.duality.dex/types"
import { TickMap } from "NicholasDotSol-duality-client-ts/nicholasdotsol.duality.dex/types"
import { TokenMap } from "NicholasDotSol-duality-client-ts/nicholasdotsol.duality.dex/types"
import { TokenPairType } from "NicholasDotSol-duality-client-ts/nicholasdotsol.duality.dex/types"
import { Tokens } from "NicholasDotSol-duality-client-ts/nicholasdotsol.duality.dex/types"


export { AdjanceyMatrix, EdgeRow, FeeList, LimitOrderTrancheTrancheIndexes, LimitOrderTranche, LimitOrderTrancheUser, PairMap, Params, Reserve0AndSharesType, Shares, TickDataType, TickMap, TokenMap, TokenPairType, Tokens };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
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

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Params: {},
				TickMap: {},
				TickMapAll: {},
				PairMap: {},
				PairMapAll: {},
				Tokens: {},
				TokensAll: {},
				TokenMap: {},
				TokenMapAll: {},
				Shares: {},
				SharesAll: {},
				FeeList: {},
				FeeListAll: {},
				EdgeRow: {},
				EdgeRowAll: {},
				AdjanceyMatrix: {},
				AdjanceyMatrixAll: {},
				LimitOrderTrancheUser: {},
				LimitOrderTrancheUserAll: {},
				LimitOrderTranche: {},
				LimitOrderTrancheAll: {},
				
				_Structure: {
						AdjanceyMatrix: getStructure(AdjanceyMatrix.fromPartial({})),
						EdgeRow: getStructure(EdgeRow.fromPartial({})),
						FeeList: getStructure(FeeList.fromPartial({})),
						LimitOrderTrancheTrancheIndexes: getStructure(LimitOrderTrancheTrancheIndexes.fromPartial({})),
						LimitOrderTranche: getStructure(LimitOrderTranche.fromPartial({})),
						LimitOrderTrancheUser: getStructure(LimitOrderTrancheUser.fromPartial({})),
						PairMap: getStructure(PairMap.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						Reserve0AndSharesType: getStructure(Reserve0AndSharesType.fromPartial({})),
						Shares: getStructure(Shares.fromPartial({})),
						TickDataType: getStructure(TickDataType.fromPartial({})),
						TickMap: getStructure(TickMap.fromPartial({})),
						TokenMap: getStructure(TokenMap.fromPartial({})),
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
				getTickMap: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TickMap[JSON.stringify(params)] ?? {}
		},
				getTickMapAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TickMapAll[JSON.stringify(params)] ?? {}
		},
				getPairMap: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PairMap[JSON.stringify(params)] ?? {}
		},
				getPairMapAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PairMapAll[JSON.stringify(params)] ?? {}
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
				getTokenMap: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TokenMap[JSON.stringify(params)] ?? {}
		},
				getTokenMapAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TokenMapAll[JSON.stringify(params)] ?? {}
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
				getEdgeRow: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.EdgeRow[JSON.stringify(params)] ?? {}
		},
				getEdgeRowAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.EdgeRowAll[JSON.stringify(params)] ?? {}
		},
				getAdjanceyMatrix: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.AdjanceyMatrix[JSON.stringify(params)] ?? {}
		},
				getAdjanceyMatrixAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.AdjanceyMatrixAll[JSON.stringify(params)] ?? {}
		},
				getLimitOrderTrancheUser: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LimitOrderTrancheUser[JSON.stringify(params)] ?? {}
		},
				getLimitOrderTrancheUserAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LimitOrderTrancheUserAll[JSON.stringify(params)] ?? {}
		},
				getLimitOrderTranche: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LimitOrderTranche[JSON.stringify(params)] ?? {}
		},
				getLimitOrderTrancheAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LimitOrderTrancheAll[JSON.stringify(params)] ?? {}
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
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTickMap({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryTickMap( key.pairId,  key.tickIndex)).data
				
					
				commit('QUERY', { query: 'TickMap', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTickMap', payload: { options: { all }, params: {...key},query }})
				return getters['getTickMap']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTickMap API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTickMapAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryTickMapAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.NicholasdotsolDualityDex.query.queryTickMapAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'TickMapAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTickMapAll', payload: { options: { all }, params: {...key},query }})
				return getters['getTickMapAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTickMapAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPairMap({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryPairMap( key.pairId)).data
				
					
				commit('QUERY', { query: 'PairMap', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPairMap', payload: { options: { all }, params: {...key},query }})
				return getters['getPairMap']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPairMap API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPairMapAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryPairMapAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.NicholasdotsolDualityDex.query.queryPairMapAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'PairMapAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPairMapAll', payload: { options: { all }, params: {...key},query }})
				return getters['getPairMapAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPairMapAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTokens({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryTokens( key.id)).data
				
					
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
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryTokensAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.NicholasdotsolDualityDex.query.queryTokensAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'TokensAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTokensAll', payload: { options: { all }, params: {...key},query }})
				return getters['getTokensAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTokensAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTokenMap({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryTokenMap( key.address)).data
				
					
				commit('QUERY', { query: 'TokenMap', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTokenMap', payload: { options: { all }, params: {...key},query }})
				return getters['getTokenMap']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTokenMap API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryTokenMapAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryTokenMapAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.NicholasdotsolDualityDex.query.queryTokenMapAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'TokenMapAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryTokenMapAll', payload: { options: { all }, params: {...key},query }})
				return getters['getTokenMapAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryTokenMapAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryShares({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryShares( key.address,  key.pairId,  key.tickIndex,  key.fee)).data
				
					
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
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.querySharesAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.NicholasdotsolDualityDex.query.querySharesAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
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
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryFeeList( key.id)).data
				
					
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
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryFeeListAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.NicholasdotsolDualityDex.query.queryFeeListAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'FeeListAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryFeeListAll', payload: { options: { all }, params: {...key},query }})
				return getters['getFeeListAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryFeeListAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryEdgeRow({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryEdgeRow( key.id)).data
				
					
				commit('QUERY', { query: 'EdgeRow', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryEdgeRow', payload: { options: { all }, params: {...key},query }})
				return getters['getEdgeRow']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryEdgeRow API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryEdgeRowAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryEdgeRowAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.NicholasdotsolDualityDex.query.queryEdgeRowAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'EdgeRowAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryEdgeRowAll', payload: { options: { all }, params: {...key},query }})
				return getters['getEdgeRowAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryEdgeRowAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAdjanceyMatrix({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryAdjanceyMatrix( key.id)).data
				
					
				commit('QUERY', { query: 'AdjanceyMatrix', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAdjanceyMatrix', payload: { options: { all }, params: {...key},query }})
				return getters['getAdjanceyMatrix']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAdjanceyMatrix API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAdjanceyMatrixAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryAdjanceyMatrixAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.NicholasdotsolDualityDex.query.queryAdjanceyMatrixAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'AdjanceyMatrixAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAdjanceyMatrixAll', payload: { options: { all }, params: {...key},query }})
				return getters['getAdjanceyMatrixAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAdjanceyMatrixAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLimitOrderTrancheUser({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryLimitOrderTrancheUser( key.pairId,  key.token,  key.tickIndex,  key.count,  key.address)).data
				
					
				commit('QUERY', { query: 'LimitOrderTrancheUser', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLimitOrderTrancheUser', payload: { options: { all }, params: {...key},query }})
				return getters['getLimitOrderTrancheUser']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLimitOrderTrancheUser API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLimitOrderTrancheUserAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryLimitOrderTrancheUserAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.NicholasdotsolDualityDex.query.queryLimitOrderTrancheUserAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'LimitOrderTrancheUserAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLimitOrderTrancheUserAll', payload: { options: { all }, params: {...key},query }})
				return getters['getLimitOrderTrancheUserAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLimitOrderTrancheUserAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLimitOrderTranche({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryLimitOrderTranche( key.pairId,  key.token,  key.tickIndex,  key.trancheIndex)).data
				
					
				commit('QUERY', { query: 'LimitOrderTranche', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLimitOrderTranche', payload: { options: { all }, params: {...key},query }})
				return getters['getLimitOrderTranche']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLimitOrderTranche API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLimitOrderTrancheAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.NicholasdotsolDualityDex.query.queryLimitOrderTrancheAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.NicholasdotsolDualityDex.query.queryLimitOrderTrancheAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'LimitOrderTrancheAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLimitOrderTrancheAll', payload: { options: { all }, params: {...key},query }})
				return getters['getLimitOrderTrancheAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLimitOrderTrancheAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgWithdrawFilledLimitOrder({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.NicholasdotsolDualityDex.tx.sendMsgWithdrawFilledLimitOrder({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdrawFilledLimitOrder:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgWithdrawFilledLimitOrder:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCancelLimitOrder({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.NicholasdotsolDualityDex.tx.sendMsgCancelLimitOrder({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancelLimitOrder:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCancelLimitOrder:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgPlaceLimitOrder({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.NicholasdotsolDualityDex.tx.sendMsgPlaceLimitOrder({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPlaceLimitOrder:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgPlaceLimitOrder:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgSwap({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.NicholasdotsolDualityDex.tx.sendMsgSwap({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSwap:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSwap:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeposit({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.NicholasdotsolDualityDex.tx.sendMsgDeposit({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeposit:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeposit:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgWithdrawl({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.NicholasdotsolDualityDex.tx.sendMsgWithdrawl({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdrawl:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgWithdrawl:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgWithdrawFilledLimitOrder({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.NicholasdotsolDualityDex.tx.msgWithdrawFilledLimitOrder({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdrawFilledLimitOrder:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgWithdrawFilledLimitOrder:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCancelLimitOrder({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.NicholasdotsolDualityDex.tx.msgCancelLimitOrder({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancelLimitOrder:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCancelLimitOrder:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgPlaceLimitOrder({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.NicholasdotsolDualityDex.tx.msgPlaceLimitOrder({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgPlaceLimitOrder:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgPlaceLimitOrder:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgSwap({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.NicholasdotsolDualityDex.tx.msgSwap({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSwap:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSwap:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeposit({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.NicholasdotsolDualityDex.tx.msgDeposit({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeposit:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeposit:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgWithdrawl({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.NicholasdotsolDualityDex.tx.msgWithdrawl({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdrawl:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgWithdrawl:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
