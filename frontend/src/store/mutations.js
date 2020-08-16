export const setVersion = (state, v) => {
  state.version = v;
};

export const setEnvironments = (state, data) => {
  state.environments = data.envList;
};

export const setEnvironment = (state, { id, data }) => {
  state.environments = state.environments.map((env) => (env.id === id ? data : env));
};

export const setEnvironmentLoading = (state, { id, flag }) => {
  state.environmentsLoading = {
    ...state.environmentsLoading,
    [id]: flag,
  };
};

export const setInstanceLoading = (state, { id, flag }) => {
  state.instancesLoading = {
    ...state.instancesLoading,
    [id]: flag,
  };
};

export const setInstanceStateStatus = (state, { id, status, envId }) => {
  state.environments = state.environments.map((env) => {
    if (env.id !== envId) {
      return env;
    }

    const newInstances = env.instances.map(
      (instance) => (instance.id !== id
        ? instance
        : {
          ...instance,
          state: status,
        }),
    );

    const numInstancesRunning = newInstances.filter((x) => x.state === 'running').length;
    const newState = numInstancesRunning === env.total_instances
      ? 'running'
      : numInstancesRunning === 0
        ? 'stopped'
        : 'mixed';

    return {
      ...env,
      instances: newInstances,
      running_instances: numInstancesRunning,
      state: newState,
    };
  });
};

export const setError = (state, err) => {
  state.error = err;
};

export const setIsLoading = (state, flag) => {
  state.isLoading = flag;
};
