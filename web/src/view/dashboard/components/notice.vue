<template>
  <el-scrollbar>
    <div
      v-for="(item, index) in notices"
      :key="index"
      class="flex items-center gap-3 py-1"
    >
      <div
        class="shrink-0 rounded-full border border-black/10 px-2 py-0.5 text-[11px] leading-4 text-black/70 dark:border-white/10 dark:text-white/70"
      >
        {{ item.typeTitle }}
      </div>
      <el-tooltip effect="light" :content="item.title" placement="top">
        <div class="min-w-0 text-xs text-black/70 dark:text-white/70 line-clamp-1">
          {{ item.title }}
        </div>
      </el-tooltip>
    </div>
  </el-scrollbar>
</template>

<script setup>
  const notices = [
    {
      type: 'success',
      typeTitle: '通知',
      title: '将进入专属飞书群，获取辅助。'
    },
    {
      type: 'warning',
      typeTitle: '警告',
      title: '警告信息。'
    },
    {
      type: 'danger',
      typeTitle: '你好',
      title: '欢迎来到DeedMX。'
    },
    {
      type: 'info',
      typeTitle: '信息',
      title: '再次感谢您'
    },
    {
      type: 'primary',
      typeTitle: '公告',
      title: '让创意更有价值。'
    },
    {
      type: 'success',
      typeTitle: '通知',
      title: '让劳动更有意义。'
    },
    {
      type: 'warning',
      typeTitle: '警告',
      title: '让思维更有深度。'
    },
    {
      type: 'danger',
      typeTitle: '错误',
      title: '让生活更有趣味。'
    },
    {
      type: 'info',
      typeTitle: '信息',
      title: '让生活更有活力。'
    }
  ]
</script>

<style scoped lang="scss"></style>
