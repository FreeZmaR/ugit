<template>
  <div id="app">
    <!-- Branch Selector -->
    <div class="branch-selector">
      <label>Select Branch:</label>
      <input type="text" v-model="searchTerm" placeholder="Search for branches..." />

      <ul class="branch-list">
        <li
            v-for="branch in filteredBranches"
            :key="branch"
            @click="selectBranch(branch)"
            :class="{ selected: branch === selectedBranch }"
        >
          <template v-if="branch !== ''">
            {{ branch }}
          </template>

        </li>
      </ul>
    </div>
    <div>
      Selected Branch: <strong>{{ selectedBranch }}</strong>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, computed } from 'vue';
import { GetBranches } from '../wailsjs/go/main/App';

export default defineComponent({
  setup() {
    const branches = ref<string[]>([]);
    const selectedBranch = ref<string>("");
    const searchTerm = ref<string>("");

    const fetchBranches = async () => {
      try {
        branches.value = await GetBranches();
      } catch (error) {
        console.error("Error fetching branches:", error);
      }
    };

    const filteredBranches = computed(() => {
      return branches.value.filter(branch =>
          branch.toLowerCase().includes(searchTerm.value.toLowerCase())
      );
    });

    const selectBranch = (branch: string) => {
      selectedBranch.value = branch;
    };

    onMounted(fetchBranches);

    return {
      branches,
      searchTerm,
      filteredBranches,
      selectBranch,
      selectedBranch,
    };
  },
});
</script>

<style scoped>
.branch-selector {
  display: flex;
  flex-direction: column;
  width: 200px;
}

.branch-list {
  list-style-type: none;
  padding: 0;
  margin-top: 8px;
  max-height: 150px;
  overflow-y: auto;
  border: 1px solid #ccc;
}

.branch-list li {
  padding: 8px;
  cursor: pointer;
}

.branch-list li.selected {
  background-color: #007bff;
  color: white;
}
</style>
