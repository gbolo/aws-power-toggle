import { shallowMount, mount } from '@vue/test-utils';
import EnvironmentList from '@/components/EnvironmentList.vue';
import Environment from '@/components/Environment.vue';

const environments = [
  {
    Name: 'env one',
    State: 'running',
  },
  {
    Name: 'env two',
    State: 'stopped',
  },
];

describe('EnvironmentList.vue', () => {
  it('renders all environments', () => {
    const wrapper = shallowMount(EnvironmentList, { propsData: { environments } });
    expect(wrapper.findAll(Environment)).toHaveLength(environments.length);
  });

  it('renders all item filters as selected by default', () => {
    const wrapper = mount(EnvironmentList, { propsData: { environments } });
    const filterList = wrapper.find('.environment-list__filter-list');
    expect(filterList.findAll('.selected')).toHaveLength(wrapper.vm.itemFilters.length);
  });
});
