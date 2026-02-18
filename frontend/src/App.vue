<template>
  <div class="min-h-screen bg-gray-950 text-white flex flex-col selection:bg-blue-500/30">
    <!-- Glassmorphism Navbar -->
    <nav class="sticky top-0 z-50 glass border-b border-white/5 py-4">
      <div class="container mx-auto px-6 flex justify-between items-center">
        <router-link to="/" class="group flex items-center gap-3">
          <div class="w-10 h-10 bg-gradient-to-br from-blue-500 to-indigo-600 rounded-xl flex items-center justify-center shadow-lg shadow-blue-500/20 group-hover:scale-110 transition-all">
            <span class="text-xl">🛒</span>
          </div>
          <span class="text-2xl font-black tracking-tighter bg-clip-text text-transparent bg-gradient-to-r from-white to-gray-400">
            COMMUNITY <span class="text-blue-500">MARKET</span>
          </span>
        </router-link>

        <div class="flex items-center gap-8">
          <div class="hidden md:flex items-center gap-6 text-sm font-medium text-gray-400">
            <router-link to="/" class="hover:text-white transition-colors">Browse</router-link>
            <template v-if="authStore.isAuthenticated">
              <router-link to="/listings/create" class="hover:text-white transition-colors">Post Item</router-link>
              <router-link to="/messages" class="hover:text-white transition-colors">Messages</router-link>
              <router-link v-if="authStore.isAdmin" to="/admin" class="hover:text-yellow-400 transition-colors">Admin</router-link>
            </template>
          </div>

          <div class="h-6 w-px bg-white/10 hidden md:block"></div>

          <div class="flex items-center gap-4">
            <template v-if="!authStore.isAuthenticated">
              <router-link to="/login" class="text-sm font-semibold hover:text-blue-400 transition-colors">Login</router-link>
              <router-link to="/register" class="px-5 py-2.5 bg-blue-600 text-white text-sm font-bold rounded-xl hover:bg-blue-500 shadow-lg shadow-blue-500/20 transition-all">
                Get Started
              </router-link>
            </template>
            <template v-else>
              <div class="flex items-center gap-3">
                <div class="text-right hidden sm:block">
                  <p class="text-xs text-gray-500 font-medium leading-none mb-1">Welcome,</p>
                  <p class="text-sm font-bold text-white leading-none">{{ authStore.user?.name }}</p>
                </div>
                <button @click="logout" class="p-2.5 bg-red-500/10 text-red-500 rounded-xl hover:bg-red-500 hover:text-white transition-all">
                  <span class="text-lg">Logout</span>
                </button>
              </div>
            </template>
          </div>
        </div>
      </div>
    </nav>

    <main class="flex-grow">
      <router-view v-slot="{ Component }">
        <transition name="fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </main>

    <footer class="bg-gray-900/50 border-t border-white/5 pt-16 pb-8">
      <div class="container mx-auto px-6">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-12 mb-12">
          <div class="space-y-4">
            <div class="flex items-center gap-2">
              <span class="text-xl">🛒</span>
              <span class="text-xl font-bold">Community Market</span>
            </div>
            <p class="text-gray-500 text-sm leading-relaxed max-w-xs">
              The premier destination for local community exchange. Buy, sell, and trade with neighbors safely.
            </p>
          </div>
          <div class="space-y-4">
            <h4 class="text-white font-bold">Quick Links</h4>
            <ul class="text-gray-500 text-sm space-y-2">
              <li><router-link to="/" class="hover:text-blue-500 transition-colors">Safety Tips</router-link></li>
              <li><router-link to="/" class="hover:text-blue-500 transition-colors">Terms of Service</router-link></li>
              <li><router-link to="/" class="hover:text-blue-500 transition-colors">Privacy Policy</router-link></li>
            </ul>
          </div>
          <div class="space-y-4">
            <h4 class="text-white font-bold">Connect</h4>
            <div class="flex gap-4">
              <a href="#" class="w-10 h-10 bg-gray-800 rounded-lg flex items-center justify-center hover:bg-blue-600 transition-all text-gray-400 hover:text-white">TW</a>
              <a href="#" class="w-10 h-10 bg-gray-800 rounded-lg flex items-center justify-center hover:bg-blue-600 transition-all text-gray-400 hover:text-white">FB</a>
              <a href="#" class="w-10 h-10 bg-gray-800 rounded-lg flex items-center justify-center hover:bg-blue-600 transition-all text-gray-400 hover:text-white">IG</a>
            </div>
          </div>
        </div>
        <div class="pt-8 border-t border-white/5 text-center">
          <p class="text-gray-600 text-sm">&copy; 2026 Community Marketplace. Built with ❤️ for the community.</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { useAuthStore } from './stores/auth';
import { useRouter } from 'vue-router';

const authStore = useAuthStore();
const router = useRouter();

const logout = () => {
  authStore.logout();
  router.push('/login');
};
</script>

<style>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
