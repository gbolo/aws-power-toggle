export const isAppLoading = (state) => state.isLoading;
export const isEnvironmentLoading = (state) => (id) => state.environmentsLoading[id] === true;
export const isInstanceLoading = (state) => (id) => state.instancesLoading[id] === true;

const isEnvironmentStateStatus = (vState, envId, status) => {
  const env = vState.environments.find((x) => x.id === envId);
  if (!env) {
    return false;
  }
  return env.state === status;
};

export const isEnvironmentStateRunning = (vState) => (id) => isEnvironmentStateStatus(vState, id, 'running');
export const isEnvironmentStateStopped = (vState) => (id) => isEnvironmentStateStatus(vState, id, 'stopped');
export const isEnvironmentStateMixed = (vState) => (id) => isEnvironmentStateStatus(vState, id, 'mixed');
export const isEnvironmentStateChanging = (vState) => (id) => isEnvironmentStateStatus(vState, id, 'changing');
