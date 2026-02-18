<template>
  <div class="py-8 max-w-6xl mx-auto px-4">
    <h1 class="text-3xl font-bold text-white mb-8">My Messages</h1>

    <div class="flex flex-col lg:flex-row gap-8 bg-gray-800 rounded-2xl border border-gray-700 overflow-hidden h-[600px]">
      <!-- Conversation List -->
      <div class="lg:w-1/3 border-r border-gray-700 overflow-y-auto">
        <div v-if="conversations.length === 0" class="p-8 text-center text-gray-500">
          No conversations yet.
        </div>
        <div v-else>
          <div v-for="conv in uniqueConversations" :key="conv.id" @click="selectThread(conv)" :class="['p-4 border-b border-gray-700 cursor-pointer transition-all hover:bg-gray-700/50', activeThread?.other_user_id === getOtherUser(conv).id ? 'bg-gray-700 border-l-4 border-l-blue-500' : '']">
            <div class="flex justify-between items-start mb-1">
              <span class="font-bold text-white">{{ getOtherUser(conv).name }}</span>
              <span class="text-xs text-gray-500">{{ new Date(conv.created_at).toLocaleDateString() }}</span>
            </div>
            <div class="text-xs text-blue-400 mb-2">Item: {{ conv.item?.title }}</div>
            <p class="text-sm text-gray-400 truncate">{{ conv.content }}</p>
          </div>
        </div>
      </div>

      <!-- Message History -->
      <div class="lg:w-2/3 flex flex-col h-full bg-gray-900/50">
        <div v-if="!activeThread" class="flex-grow flex items-center justify-center text-gray-500">
          Select a conversation to start chatting.
        </div>
        <template v-else>
          <div class="p-4 bg-gray-800 border-b border-gray-700 flex justify-between items-center">
            <div>
              <div class="font-bold text-white text-lg">{{ activeThread.other_user_name }}</div>
              <div class="text-xs text-gray-400">Discussing: {{ activeThread.item_title }}</div>
            </div>
          </div>

          <div class="flex-grow overflow-y-auto p-6 space-y-4" ref="messageBox">
            <div v-for="msg in thread" :key="msg.id" :class="['flex', msg.sender_id === authStore.user.id ? 'justify-end' : 'justify-start']">
              <div :class="['max-w-[80%] p-3 rounded-2xl text-sm', msg.sender_id === authStore.user.id ? 'bg-blue-600 text-white rounded-tr-none' : 'bg-gray-700 text-gray-200 rounded-tl-none']">
                {{ msg.content }}
                <div class="text-[10px] mt-1 opacity-50">{{ new Date(msg.created_at).toLocaleTimeString() }}</div>
              </div>
            </div>
          </div>

          <div class="p-4 bg-gray-800 border-t border-gray-700">
            <div class="flex gap-2">
              <input v-model="replyContent" @keyup.enter="sendReply" placeholder="Type your reply..." class="flex-grow bg-gray-900 border border-gray-700 rounded-full px-6 py-3 text-white focus:outline-none focus:ring-2 focus:ring-blue-500" />
              <button @click="sendReply" :disabled="!replyContent" class="bg-blue-600 p-3 rounded-full hover:bg-blue-700 transition-all disabled:opacity-50">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
                </svg>
              </button>
            </div>
          </div>
        </template>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, nextTick } from 'vue';
import { useAuthStore } from '../stores/auth';
import api from '../services/api';

const authStore = useAuthStore();
const conversations = ref([]);
const thread = ref([]);
const activeThread = ref(null);
const replyContent = ref('');
const messageBox = ref(null);

const fetchConversations = async () => {
  const res = await api.get('/messages');
  conversations.value = res.data;
};

const uniqueConversations = computed(() => {
  // Simple grouping by other user and item for the list
  const seen = new Set();
  return conversations.value.filter(conv => {
    const otherUser = getOtherUser(conv);
    const key = `${otherUser.id}-${conv.item_id}`;
    if (seen.has(key)) return false;
    seen.add(key);
    return true;
  });
});

const getOtherUser = (conv) => {
  return conv.sender_id === authStore.user.id ? conv.receiver : conv.sender;
};

const selectThread = async (conv) => {
  const otherUser = getOtherUser(conv);
  activeThread.value = {
    other_user_id: otherUser.id,
    other_user_name: otherUser.name,
    item_id: conv.item_id,
    item_title: conv.item?.title
  };
  await fetchThread();
  scrollToBottom();
};

const fetchThread = async () => {
  if (!activeThread.value) return;
  const res = await api.get(`/messages/${activeThread.value.other_user_id}`, {
    params: { item_id: activeThread.value.item_id }
  });
  thread.value = res.data;
};

const sendReply = async () => {
  if (!replyContent.value) return;
  try {
    await api.post('/messages', {
      receiver_id: activeThread.value.other_user_id,
      item_id: activeThread.value.item_id,
      content: replyContent.value
    });
    replyContent.value = '';
    await fetchThread();
    scrollToBottom();
  } catch (error) {
    alert('Failed to send reply');
  }
};

const scrollToBottom = () => {
  nextTick(() => {
    if (messageBox.value) {
      messageBox.value.scrollTop = messageBox.value.scrollHeight;
    }
  });
};

onMounted(fetchConversations);
</script>
