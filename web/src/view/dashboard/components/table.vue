<template>
  <div>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="ranking" label="排名" width="80" align="center" />
      <el-table-column prop="message" label="更新内容" show-overflow-tooltip />
      <el-table-column prop="author" label="提交人" width="140" />
      <el-table-column prop="date" label="时间" width="180" />
    </el-table>
  </div>
</template>

<script setup>
  // const tableData = [
  //   {
  //     ranking: 1,
  //     title: '更简洁的使用，更快速的操作',
  //     click_num: 523,
  //     hot: 263
  //   },
  //   {
  //     ranking: 2,
  //     title: '更优质的服务，更便捷的使用体验',
  //     click_num: 416,
  //     hot: 223
  //   },
  //   {
  //     ranking: 3,
  //     title: '更快速的创意实现，更高效的工作效率',
  //     click_num: 337,
  //     hot: 176
  //   },
  //   {
  //     ranking: 4,
  //     title: '更多的创意资源，更多的创意灵感',
  //     click_num: 292,
  //     hot: 145
  //   },
  //   {
  //     ranking: 5,
  //     title: '更合理的结构，更清晰的逻辑',
  //     click_num: 173,
  //     hot: 110
  //   }
  // ]

  import { Commits } from '@/api/github'
  import { formatTimeToStr } from '@/utils/date'
  import { ref, onMounted } from 'vue'

  const tableData = ref([])

  const loadCommits = async () => {
    const { data } = await Commits(1)
    tableData.value = data.slice(0, 5).map((item, index) => {
      return {
        ranking: index + 1,
        message: item.commit.message,
        author: item.commit.author.name,
        date: formatTimeToStr(item.commit.author.date, 'yyyy-MM-dd hh:mm:ss')
      }
    })
  }

  onMounted(() => {
    loadCommits()
  })

</script>

<style scoped lang="scss"></style>
