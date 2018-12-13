export const setVersion = (state, v) => {
  state.version = v;
};

export const setEnvironments = (state, envs) => {
  state.environments = envs;
};

export const setEnvironment = (state, { id, data }) => {
  state.environments = state.environments.map(env => (env.id === id ? data : env));
};

export const setError = (state, err) => {
  state.error = err;
};

export const setEnvironmentLoading = (state, { id, flag }) => {
  state.environmentsLoading = {
    ...state.environmentsLoading,
    [id]: flag,
  };
};
