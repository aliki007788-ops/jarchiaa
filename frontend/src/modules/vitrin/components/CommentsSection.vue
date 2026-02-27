<template>
  <div class="comments-section">
    <h3>نظرات ({{ comments.length }})</h3>

    <!-- فرم ارسال نظر -->
    <div class="add-comment" v-if="isLoggedIn">
      <textarea 
        v-model="newCommentText" 
        placeholder="نظرت رو بنویس..."
        rows="3"
      ></textarea>
      <button @click="submitComment" :disabled="!newCommentText">
        ارسال نظر
      </button>
    </div>
    <div v-else class="login-to-comment">
      <p>برای ارسال نظر <span @click="goToLogin">وارد شوید</span></p>
    </div>

    <!-- لیست نظرات -->
    <div class="comments-list">
      <div v-for="comment in comments" :key="comment.id" class="comment-item">
        <div class="comment-header">
          <span class="comment-author">{{ comment.user_name }}</span>
          <span class="comment-time">{{ comment.time }}</span>
        </div>
        <p class="comment-text">{{ comment.text }}</p>
        <div class="comment-actions">
          <button @click="likeComment(comment.id)">
            <i :class="comment.isLiked ? 'fas fa-heart' : 'far fa-heart'"></i>
            {{ comment.likes }}
          </button>
          <button @click="replyToComment(comment.id)">پاسخ</button>
        </div>

        <!-- پاسخ‌ها (اگر داشت) -->
        <div v-if="comment.replies && comment.replies.length" class="replies">
          <div v-for="reply in comment.replies" :key="reply.id" class="reply-item">
            <!-- نمایش پاسخ -->
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'

export default {
  props: {
    adId: { type: String, required: true },
    comments: { type: Array, default: () => [] }
  },
  emits: ['newComment'],
  setup(props, { emit }) {
    const router = useRouter()
    const newCommentText = ref('')
    const isLoggedIn = computed(() => !!localStorage.getItem('token'))

    const submitComment = () => {
      if (newCommentText.value.trim()) {
        emit('newComment', newCommentText.value)
        newCommentText.value = ''
      }
    }

    const goToLogin = () => {
      router.push('/login')
    }

    const likeComment = (commentId) => {
      // TODO: API لایک نظر
    }

    const replyToComment = (commentId) => {
      // TODO: پاسخ به نظر
    }

    return {
      newCommentText,
      isLoggedIn,
      submitComment,
      goToLogin,
      likeComment,
      replyToComment
    }
  }
}
</script>