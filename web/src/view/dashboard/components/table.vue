<template>
  <div>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="ranking" label="排名" width="80" align="center" />
      <el-table-column prop="title" label="标题" show-overflow-tooltip />
      <el-table-column prop="click_num" label="点击量" width="140" />
      <el-table-column prop="hot" label="热度" width="180" />
    </el-table>
  </div>
</template>

<script setup>
  const tableData = [
    {
      ranking: 1,
      title: '更简洁的使用，更快速的操作',
      click_num: 523,
      hot: 263
    },
    {
      ranking: 2,
      title: '更优质的服务，更便捷的使用体验',
      click_num: 416,
      hot: 223
    },
    {
      ranking: 3,
      title: '更快速的创意实现，更高效的工作效率',
      click_num: 337,
      hot: 176
    },
    {
      ranking: 4,
      title: '更多的创意资源，更多的创意灵感',
      click_num: 292,
      hot: 145
    },
    {
      ranking: 5,
      title: '更合理的结构，更清晰的逻辑',
      click_num: 173,
      hot: 110
    }
  ]


</script>

<style scoped lang="scss"></style>
