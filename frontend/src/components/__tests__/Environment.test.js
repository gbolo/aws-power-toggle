import { shallowMount } from '@vue/test-utils';
import Environment from '@/components/Environment.vue';

describe('Environment.vue', () => {
  it('renders the name of the environment passed', () => {
    const env = {
      Name: 'Environment Name',
    };
    const wrapper = shallowMount(Environment, { propsData: { env } });
    expect(wrapper.text()).toMatch(env.Name);
  });
});
