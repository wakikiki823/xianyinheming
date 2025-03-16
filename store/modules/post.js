export default {
  state: {
    posts: [],
    drafts: []
  },
  mutations: {
    ADD_POST(state, post) {
      state.posts.unshift(post);
    },
    SAVE_DRAFT(state, draft) {
      state.drafts.push(draft);
    }
  },
  actions: {
    async publishPost({ commit }, postData) {
      try {
        const res = await api.publishPost(postData);
        commit('ADD_POST', res.data);
        return res;
      } catch (err) {
        throw err;
      }
    }
  }
} 