export const setVersion = (state, v) => {
  state.version = v;
};

export const setEnvironments = (state, data) => {
  state.environments = data.envList;
  state.totalBillsAccrued = data.totalBillsAccrued;
  state.totalBillsSaved = data.totalBillsSaved;
};

export const setEnvironment = (state, { id, data }) => {
  state.environments = state.environments.map(env => (env.id === id ? data : env));
};

export const setEnvironmentLoading = (state, { id, flag }) => {
  state.environmentsLoading = {
    ...state.environmentsLoading,
    [id]: flag,
  };
};

export const setError = (state, err) => {
  state.error = err;
};

export const setIsLoading = (state, flag) => {
  state.isLoading = flag;
};
