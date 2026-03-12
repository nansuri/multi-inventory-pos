<script setup lang="ts">
import { ref, onMounted } from 'vue';
import DashboardLayout from '../layouts/DashboardLayout.vue';
import api from '../api/axios';
import { useConfigStore } from '../stores/config';
import { Plus, UtensilsCrossed, BookOpen, Trash2, ArrowRight } from 'lucide-vue-next';

const products = ref<any[]>([]); 
const inventoryItems = ref<any[]>([]);
const loading = ref(true);
const configStore = useConfigStore();
const isAddModalOpen = ref(false);
const isRecipeModalOpen = ref(false);
const selectedProduct = ref<any>(null);

const newProduct = ref({
  name: '',
  price: 0
});

const currentRecipe = ref<any[]>([]);

const fetchProducts = async () => {
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

const addProduct = async () => {
  try {
    await api.post('/api/products', newProduct.value);
    isAddModalOpen.value = false;
    newProduct.value = { name: '', price: 0 };
    await fetchProducts();
  } catch (error) {
    console.error('Failed to add recipe', error);
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
  try {
    await api.post(`/api/products/${selectedProduct.value.id}/recipe`, {
      recipes: currentRecipe.value
    });
    isRecipeModalOpen.value = false;
    await fetchProducts();
  } catch (error) {
    console.error('Failed to save recipe', error);
  }
};

onMounted(() => {
  fetchProducts();
  fetchInventory();
});
</script>

<template>
  <DashboardLayout>
    <div class="space-y-8">
      <!-- Header -->
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
        <div>
          <h1 class="text-2xl font-black text-slate-800 tracking-tight flex items-center gap-3">
            <BookOpen class="w-8 h-8 text-indigo-600" />
            {{ $t('recipes.blueprints') }}
          </h1>
          <p class="text-slate-500 font-medium mt-1">Define the formulas for your menu items.</p>
        </div>
        <button 
          @click="isAddModalOpen = true"
          class="bg-indigo-600 text-white px-8 py-4 rounded-2xl font-black shadow-lg shadow-indigo-100 hover:bg-indigo-700 transition-all flex items-center gap-2"
        >
          <Plus class="w-5 h-5" />
          {{ $t('recipes.createBlueprint') }}
        </button>
      </div>

      <!-- Recipe Cards -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <div 
          v-for="product in products" 
          :key="product.id" 
          class="bg-white rounded-[2.5rem] shadow-sm border border-slate-100 p-8 flex flex-col group hover:shadow-xl transition-all duration-300 relative overflow-hidden"
        >
          <div class="absolute top-0 right-0 p-10 opacity-5 group-hover:scale-110 transition-transform">
            <UtensilsCrossed class="w-32 h-32 text-indigo-600" />
          </div>

          <div class="flex items-center gap-4 mb-8">
            <div class="p-4 bg-indigo-50 rounded-2xl text-indigo-600 group-hover:bg-indigo-600 group-hover:text-white transition-all">
              <UtensilsCrossed class="w-6 h-6" />
            </div>
            <div>
              <h3 class="text-xl font-black text-slate-800">{{ product.name }}</h3>
              <p class="text-xs font-bold text-slate-400 uppercase tracking-widest">{{ configStore.formatCurrency(product.price) }} Unit Price</p>
            </div>
          </div>
          
          <div class="space-y-4 mb-10 flex-1 relative z-10">
            <p class="text-[10px] font-black text-slate-400 uppercase tracking-[0.2em] mb-4">{{ $t('recipes.formula') }}</p>
            <div v-if="product.recipes && product.recipes.length" class="space-y-3">
              <div v-for="r in product.recipes" :key="r.id" class="flex items-center justify-between group/item">
                <span class="text-sm font-bold text-slate-600 group-hover/item:text-slate-900 transition-colors">{{ r.ingredient?.name }}</span>
                <span class="text-xs font-black text-slate-400 bg-slate-50 px-2 py-1 rounded-lg">{{ r.quantity }} {{ r.unit }}</span>
              </div>
            </div>
            <div v-else class="py-10 text-center border-2 border-dashed border-slate-100 rounded-3xl">
              <p class="text-xs font-bold text-slate-400 italic">No formula defined yet.</p>
            </div>
          </div>

          <button 
            @click="openRecipeEditor(product)" 
            class="w-full py-4 bg-slate-50 text-slate-700 rounded-2xl font-black hover:bg-indigo-600 hover:text-white transition-all flex items-center justify-center gap-2 group/btn"
          >
            {{ $t('recipes.editFormula') }}
            <ArrowRight class="w-4 h-4 group-hover/btn:translate-x-1 transition-transform" />
          </button>
        </div>

        <div v-if="products.length === 0 && !loading" class="col-span-full py-32 text-center bg-white rounded-[3rem] border-4 border-dashed border-slate-100">
          <div class="max-w-xs mx-auto space-y-4 opacity-40">
            <BookOpen class="w-20 h-20 mx-auto text-slate-400" />
            <p class="text-xl font-black text-slate-500 uppercase">{{ $t('common.noData') }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Modals (Add & Edit) -->
    <!-- Add Recipe Modal -->
    <div v-if="isAddModalOpen" class="fixed inset-0 bg-slate-900/40 backdrop-blur-sm flex items-center justify-center p-4 z-50 animate-in fade-in duration-200">
      <div class="bg-white rounded-[2.5rem] w-full max-w-md p-10 shadow-2xl scale-in-center">
        <h3 class="text-2xl font-black text-slate-800 mb-2">{{ $t('recipes.newBlueprint') }}</h3>
        <p class="text-sm font-medium text-slate-400 mb-8">Set the foundation for a new menu item.</p>
        
        <form @submit.prevent="addProduct" class="space-y-6">
          <div class="space-y-2">
            <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('recipes.dishName') }}</label>
            <input v-model="newProduct.name" type="text" required class="w-full px-5 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-bold focus:bg-white focus:border-indigo-600 transition-all outline-none" />
          </div>
          <div class="space-y-2">
            <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('recipes.sellingPrice') }}</label>
            <input v-model.number="newProduct.price" type="number" step="0.01" required class="w-full px-5 py-4 bg-slate-50 border-2 border-transparent rounded-2xl font-bold focus:bg-white focus:border-indigo-600 transition-all outline-none" />
          </div>
          
          <div class="flex flex-col gap-2 pt-4">
            <button type="submit" class="w-full py-5 bg-indigo-600 text-white rounded-2xl font-black shadow-lg shadow-indigo-100 hover:bg-indigo-700 transition-all">
              {{ $t('common.save') }}
            </button>
            <button type="button" @click="isAddModalOpen = false" class="w-full py-4 text-slate-400 font-bold hover:text-slate-600 transition-colors">
              {{ $t('common.cancel') }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Edit Formula Modal -->
    <div v-if="isRecipeModalOpen" class="fixed inset-0 bg-slate-900/40 backdrop-blur-sm flex items-center justify-center p-4 z-50 animate-in fade-in duration-200">
      <div class="bg-white rounded-[3rem] w-full max-w-2xl p-10 shadow-2xl scale-in-center overflow-hidden flex flex-col max-h-[90vh]">
        <div class="flex items-center gap-4 mb-8">
          <div class="p-4 bg-indigo-100 text-indigo-600 rounded-2xl">
            <UtensilsCrossed class="w-8 h-8" />
          </div>
          <div>
            <h3 class="text-2xl font-black text-slate-800">{{ selectedProduct?.name }}</h3>
            <p class="text-sm font-medium text-slate-400 uppercase tracking-widest">{{ $t('recipes.formula') }}</p>
          </div>
        </div>
        
        <div class="flex-1 overflow-auto pr-2 space-y-4 mb-8">
          <div v-for="(r, index) in currentRecipe" :key="index" class="grid grid-cols-12 gap-4 items-end bg-slate-50 p-5 rounded-[1.5rem] border border-slate-100 group/row">
            <div class="col-span-6 space-y-1">
              <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('inventory.ingredientName') }}</label>
              <select v-model="r.ingredient_id" class="w-full px-4 py-3 bg-white border-2 border-transparent rounded-xl font-bold focus:border-indigo-600 outline-none appearance-none transition-all">
                <option v-for="item in inventoryItems" :key="item.id" :value="item.id">{{ item.name }}</option>
              </select>
            </div>
            <div class="col-span-2 space-y-1">
              <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Qty</label>
              <input v-model.number="r.quantity" type="number" step="0.1" class="w-full px-4 py-3 bg-white border-2 border-transparent rounded-xl font-black text-center focus:border-indigo-600 outline-none transition-all" />
            </div>
            <div class="col-span-3 space-y-1">
              <label class="block text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">{{ $t('common.unit') }}</label>
              <select v-model="r.unit" class="w-full px-4 py-3 bg-white border-2 border-transparent rounded-xl font-bold focus:border-indigo-600 outline-none appearance-none transition-all">
                <option>KG</option><option>GR</option><option>Piece</option><option>Litre</option><option>ML</option>
              </select>
            </div>
            <div class="col-span-1 pb-1">
              <button @click="removeIngredientFromRecipe(index)" class="p-3 text-red-400 hover:bg-red-50 hover:text-red-600 rounded-xl transition-all">
                <Trash2 class="w-5 h-5" />
              </button>
            </div>
          </div>
          
          <button @click="addIngredientToRecipe" class="w-full py-6 border-4 border-dashed border-slate-100 rounded-[1.5rem] text-slate-400 font-black hover:border-indigo-200 hover:text-indigo-600 hover:bg-indigo-50 transition-all flex items-center justify-center gap-3">
            <Plus class="w-6 h-6" />
            {{ $t('recipes.addRequirement') }}
          </button>
        </div>

        <div class="flex gap-4 pt-6 border-t border-slate-100">
          <button type="button" @click="isRecipeModalOpen = false" class="flex-1 py-4 text-slate-400 font-bold hover:text-slate-600 transition-colors">{{ $t('common.cancel') }}</button>
          <button @click="saveRecipe" class="flex-[2] py-4 bg-indigo-600 text-white rounded-2xl font-black shadow-lg shadow-indigo-100 hover:bg-indigo-700 transition-all">
            {{ $t('common.save') }}
          </button>
        </div>
      </div>
    </div>
  </DashboardLayout>
</template>
