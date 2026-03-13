<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import ConfirmModal from '../components/ConfirmModal.vue';
import FormModal from '../components/FormModal.vue';
import ViewToggle from '../components/ViewToggle.vue';
import api from '../api/axios';
import { useConfigStore } from '../stores/config';
import { useAuthStore } from '../stores/auth';
import { useI18n } from 'vue-i18n';
import { Plus, UtensilsCrossed, BookOpen, Trash2, ArrowRight, Edit2 } from 'lucide-vue-next';

const { t } = useI18n();
const authStore = useAuthStore();
const products = ref<any[]>([]); 
const inventoryItems = ref<any[]>([]);
const loading = ref(true);
const actionLoading = ref(false);
const configStore = useConfigStore();
const viewMode = ref<'list' | 'grid'>('list'); // Default to list

const isAddModalOpen = ref(false);
const isEditModalOpen = ref(false);
const isRecipeModalOpen = ref(false);
const isDeleteModalOpen = ref(false);

const selectedProduct = ref<any>(null);

const productForm = ref({
  id: 0,
  name: '',
  price: 0
});

const currentRecipe = ref<any[]>([]);

const fetchProducts = async () => {
  loading.value = true;
  try {
    const response = await api.get('/api/products');
    products.value = response.data;
  } catch (error) {
    console.error('Failed to fetch recipes', error);
  } finally {
    loading.value = false;
  }
};

const fetchInventory = async () => {
  try {
    const response = await api.get('/api/inventory/items');
    inventoryItems.value = response.data;
  } catch (error) {
    console.error('Failed to fetch inventory', error);
  }
};

const createProduct = async () => {
  actionLoading.value = true;
  try {
    await api.post('/api/products', productForm.value);
    isAddModalOpen.value = false;
    productForm.value = { id: 0, name: '', price: 0 };
    await fetchProducts();
  } finally {
    actionLoading.value = false;
  }
};

const openEditModal = (product: any) => {
  productForm.value = { id: product.id, name: product.name, price: product.price };
  isEditModalOpen.value = true;
};

const updateProduct = async () => {
  actionLoading.value = true;
  try {
    await api.put(`/api/products/${productForm.value.id}`, productForm.value);
    isEditModalOpen.value = false;
    await fetchProducts();
  } finally {
    actionLoading.value = false;
  }
};

const deleteProduct = async () => {
  actionLoading.value = true;
  try {
    await api.delete(`/api/products/${selectedProduct.value.id}`);
    isDeleteModalOpen.value = false;
    await fetchProducts();
  } finally {
    actionLoading.value = false;
  }
};

const openRecipeEditor = (product: any) => {
  selectedProduct.value = product;
  currentRecipe.value = (product.recipes || []).map((r: any) => ({
    ingredient_id: r.ingredient_id,
    quantity: r.quantity,
    unit: r.unit
  }));
  isRecipeModalOpen.value = true;
};

const addIngredientToRecipe = () => {
  currentRecipe.value.push({ ingredient_id: null, quantity: 1, unit: 'GR' });
};

const removeIngredientFromRecipe = (index: number) => {
  currentRecipe.value.splice(index, 1);
};

const saveRecipe = async () => {
  actionLoading.value = true;
  try {
    await api.post(`/api/products/${selectedProduct.value.id}/recipe`, {
      recipes: currentRecipe.value
    });
    isRecipeModalOpen.value = false;
    await fetchProducts();
  } finally {
    actionLoading.value = false;
  }
};

onMounted(() => {
  fetchProducts();
  fetchInventory();
});

// Refresh data when branch changes
watch(() => authStore.selectedBranchID, () => {
  fetchProducts();
  fetchInventory();
});
</script>

<template>
  <div class="space-y-6 animate-in fade-in duration-500">
    <!-- Header -->
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
      <div class="flex flex-1 gap-4 items-center">
        <h1 class="text-xl font-black text-slate-800 dark:text-slate-100 tracking-tight flex items-center gap-2">
          <BookOpen class="w-6 h-6 text-indigo-600" />
          {{ t('recipes.blueprints') }}
        </h1>
        <ViewToggle v-model:mode="viewMode" />
      </div>
      <button 
        @click="isAddModalOpen = true; productForm = { id: 0, name: '', price: 0 }"
        class="bg-indigo-600 text-white px-6 py-2.5 rounded-xl text-sm font-black shadow-lg shadow-indigo-100 dark:shadow-indigo-900/20 hover:bg-indigo-700 transition-all flex items-center justify-center gap-2"
      >
        <Plus class="w-4 h-4" />
        {{ t('recipes.createBlueprint') }}
      </button>
    </div>

    <!-- List View -->
    <div v-if="viewMode === 'list'" class="bg-white dark:bg-slate-900 rounded-3xl shadow-sm border border-slate-200 dark:border-slate-800 overflow-hidden text-sm">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="bg-slate-50/50 dark:bg-slate-800/50 border-b border-slate-100 dark:border-slate-800">
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ t('recipes.dishName') }}</th>
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-right">{{ t('recipes.targetPrice') }}</th>
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-center">{{ t('recipes.ingredients') }}</th>
            <th class="px-6 py-4 text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest text-right">{{ t('common.actions') }}</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-slate-50 dark:divide-slate-800">
          <tr v-for="product in products" :key="product.id" class="group hover:bg-slate-50/30 dark:hover:bg-slate-800/30 transition-colors">
            <td class="px-6 py-4 font-bold text-slate-800 dark:text-slate-100 text-sm">{{ product.name }}</td>
            <td class="px-6 py-4 text-right font-black text-slate-700 dark:text-slate-300 text-xs">{{ configStore.formatCurrency(product.price) }}</td>
            <td class="px-6 py-4 text-center">
              <span class="px-2 py-0.5 rounded-md bg-slate-100 dark:bg-slate-800 text-slate-600 dark:text-slate-400 text-[9px] font-black uppercase border border-slate-200 dark:border-slate-700">
                {{ (product.recipes || []).length }} {{ t('recipes.ingredients') }}
              </span>
            </td>
            <td class="px-6 py-4 text-right">
              <div class="flex justify-end gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
                <button @click="openRecipeEditor(product)" class="px-2 py-1 text-indigo-600 hover:bg-indigo-50 dark:hover:bg-indigo-900/30 rounded-lg transition-all font-bold text-[10px] uppercase tracking-wider">{{ t('recipes.configureFormula') }}</button>
                <button @click="openEditModal(product)" class="p-1.5 text-slate-400 hover:text-blue-600 transition-all"><Edit2 class="w-4 h-4" /></button>
                <button @click="selectedProduct = product; isDeleteModalOpen = true" class="p-1.5 text-slate-400 hover:text-red-600 transition-all"><Trash2 class="w-4 h-4" /></button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Grid View -->
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
      <div v-for="product in products" :key="product.id" class="bg-white dark:bg-slate-900 rounded-[2rem] shadow-sm border border-slate-200 dark:border-slate-800 p-6 flex flex-col group hover:shadow-lg transition-all duration-300 relative overflow-hidden">
        <div class="absolute top-0 right-0 p-6 opacity-5 group-hover:scale-110 transition-transform">
          <UtensilsCrossed class="w-24 h-24 text-indigo-600" />
        </div>
        
        <div class="flex justify-between items-start mb-6 relative z-10">
          <div class="flex items-center gap-3">
            <div class="p-3 bg-indigo-50 dark:bg-indigo-900/30 rounded-xl text-indigo-600 dark:text-indigo-400 group-hover:bg-indigo-600 group-hover:text-white transition-all">
              <UtensilsCrossed class="w-5 h-5" />
            </div>
            <div>
              <h3 class="text-sm font-black text-slate-800 dark:text-slate-100">{{ product.name }}</h3>
              <p class="text-[9px] font-bold text-slate-400 dark:text-slate-500 uppercase tracking-widest">{{ configStore.formatCurrency(product.price) }}</p>
            </div>
          </div>
          <div class="flex gap-0.5 opacity-0 group-hover:opacity-100 transition-opacity relative z-20">
            <button @click="openEditModal(product)" class="p-1.5 text-slate-400 dark:text-slate-500 hover:text-blue-600 dark:hover:text-blue-400 transition-all"><Edit2 class="w-3.5 h-3.5" /></button>
            <button @click="selectedProduct = product; isDeleteModalOpen = true" class="p-1.5 text-slate-400 dark:text-slate-500 hover:text-red-600 dark:hover:text-red-400 transition-all"><Trash2 class="w-3.5 h-3.5" /></button>
          </div>
        </div>

        <div class="space-y-2 mb-6 flex-1 relative z-10">
          <p class="text-[8px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-[0.2em] mb-2">{{ t('recipes.formula') }}</p>
          <div v-if="product.recipes && product.recipes.length" class="space-y-1.5">
            <div v-for="r in product.recipes.slice(0, 3)" :key="r.id" class="flex items-center justify-between">
              <span class="text-[11px] font-bold text-slate-600 dark:text-slate-400 truncate max-w-[100px]">{{ r.ingredient?.name }}</span>
              <span class="text-[9px] font-black text-slate-400 dark:text-slate-500">{{ r.quantity }}{{ r.unit }}</span>
            </div>
            <p v-if="product.recipes.length > 3" class="text-[8px] text-slate-400 italic mt-1">+ {{ product.recipes.length - 3 }} more ingredients</p>
          </div>
          <div v-else class="py-4 text-center border border-dashed border-slate-100 dark:border-slate-800 rounded-xl">
            <p class="text-[9px] font-bold text-slate-400 italic">{{ t('recipes.noFormula') }}</p>
          </div>
        </div>

        <button @click="openRecipeEditor(product)" class="w-full py-2.5 bg-slate-50 dark:bg-slate-800 text-slate-700 dark:text-slate-300 rounded-xl font-black text-[10px] uppercase tracking-widest hover:bg-indigo-600 hover:text-white transition-all flex items-center justify-center gap-2 group/btn relative z-10 shadow-sm">
          {{ t('recipes.configureFormula') }}
          <ArrowRight class="w-3.5 h-3.5 group-hover/btn:translate-x-0.5 transition-transform" />
        </button>
      </div>
    </div>

    <!-- Empty State -->
    <div v-if="products.length === 0 && !loading" class="py-20 text-center bg-white dark:bg-slate-900 rounded-[2.5rem] border-2 border-dashed border-slate-200 dark:border-slate-800">
      <div class="max-w-xs mx-auto space-y-3 opacity-40 dark:opacity-20">
        <BookOpen class="w-16 h-16 mx-auto text-slate-400 dark:text-slate-600" />
        <p class="text-sm font-black text-slate-500 dark:text-slate-400 uppercase tracking-widest">{{ t('common.noData') }}</p>
      </div>
    </div>

    <!-- Modals -->
    <FormModal :show="isAddModalOpen || isEditModalOpen" :title="isEditModalOpen ? t('common.edit') + ' ' + t('recipes.blueprints') : t('recipes.newBlueprint')" :icon="BookOpen" :loading="actionLoading" @submit="isEditModalOpen ? updateProduct() : createProduct()" @cancel="isAddModalOpen = false; isEditModalOpen = false">
      <div class="space-y-4">
        <div class="space-y-1.5"><label class="block text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-1">{{ t('recipes.dishName') }}</label><input v-model="productForm.name" type="text" required class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none dark:text-slate-100 text-sm" /></div>
        <div class="space-y-1.5"><label class="block text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-1">{{ t('recipes.sellingPrice') }}</label><input v-model.number="productForm.price" type="number" step="0.01" required class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl font-bold focus:bg-white dark:focus:bg-slate-900 focus:border-indigo-600 transition-all outline-none dark:text-slate-100 text-sm" /></div>
      </div>
    </FormModal>

    <FormModal :show="isRecipeModalOpen" :title="selectedProduct?.name" :subtitle="t('recipes.formula')" :icon="UtensilsCrossed" :loading="actionLoading" maxWidth="max-w-3xl" @submit="saveRecipe" @cancel="isRecipeModalOpen = false">
      <div class="space-y-3 max-h-[50vh] overflow-auto pr-2 custom-scrollbar">
        <div v-for="(r, index) in currentRecipe" :key="index" class="grid grid-cols-12 gap-3 items-end bg-slate-50 dark:bg-slate-800/50 p-4 rounded-2xl border border-slate-100 dark:border-slate-700 group/row">
          <div class="col-span-6 space-y-1">
            <label class="block text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-1">{{ t('inventory.ingredientName') }}</label>
            <select v-model="r.ingredient_id" class="w-full px-3 py-2 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-700 rounded-lg font-bold focus:border-indigo-600 outline-none appearance-none transition-all dark:text-slate-100 text-xs">
              <option v-for="item in inventoryItems" :key="item.id" :value="item.id">{{ item.name }}</option>
            </select>
          </div>
          <div class="col-span-2 space-y-1">
            <label class="block text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-1">{{ t('common.qty') }}</label>
            <input v-model.number="r.quantity" type="number" step="0.1" class="w-full px-3 py-2 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-700 rounded-lg font-black text-center focus:border-indigo-600 outline-none transition-all dark:text-slate-100 text-xs" />
          </div>
          <div class="col-span-3 space-y-1">
            <label class="block text-[9px] font-black text-slate-400 dark:text-slate-500 uppercase tracking-widest ml-1">{{ t('common.unit') }}</label>
            <select v-model="r.unit" class="w-full px-3 py-2 bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-700 rounded-lg font-bold focus:border-indigo-600 outline-none appearance-none transition-all dark:text-slate-100 text-xs">
              <option>KG</option><option>GR</option><option>Piece</option><option>Litre</option><option>ML</option>
            </select>
          </div>
          <div class="col-span-1 pb-0.5 flex justify-center">
            <button type="button" @click="removeIngredientFromRecipe(index)" class="p-2 text-red-400 hover:text-red-600 rounded-lg transition-all"><Trash2 class="w-4 h-4" /></button>
          </div>
        </div>
        
        <button 
          type="button" 
          @click="addIngredientToRecipe" 
          class="w-full py-6 border-2 border-dashed border-slate-100 dark:border-slate-800 rounded-2xl text-slate-400 dark:text-slate-600 font-black hover:border-indigo-200 dark:hover:border-indigo-900 hover:text-indigo-600 dark:hover:text-indigo-400 hover:bg-indigo-50 dark:hover:bg-indigo-950/20 transition-all flex items-center justify-center gap-2 text-[10px] uppercase tracking-widest"
        >
          <Plus class="w-4 h-4" />
          {{ t('recipes.addRequirement') }}
        </button>
      </div>
    </FormModal>

    <ConfirmModal :show="isDeleteModalOpen" type="danger" :title="t('common.confirm')" :message="`${t('common.confirm')} delete ${selectedProduct?.name}?`" :loading="actionLoading" @confirm="deleteProduct" @cancel="isDeleteModalOpen = false" />
  </div>
</template>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: var(--color-slate-200);
  border-radius: 10px;
}
.dark .custom-scrollbar::-webkit-scrollbar-thumb {
  background: var(--color-slate-800);
}
</style>
