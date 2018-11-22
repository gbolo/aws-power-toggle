import { shallowMount } from '@vue/test-utils';
import Instance from '@/components/Instance.vue';

describe('Instance.vue', () => {
  it('renders props.name when passed', () => {
    const name = 'Instance Name';
    const wrapper = shallowMount(Instance, {
      propsData: { name },
    });
    expect(wrapper.text()).toMatch(msg);
  });
});
