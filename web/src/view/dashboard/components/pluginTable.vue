<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <div>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="ranking" label="排名" width="80" align="center" />
      <el-table-column prop="title" label="插件标题" show-overflow-tooltip>
        <template #default="{ row }">
          <a class="text-active" :href="row.link" target="_blank">{{
            row.title
          }}</a>
        </template>
      </el-table-column>
      <el-table-column prop="click_num" label="关注度" width="100" />
      <el-table-column prop="hot" label="热度值" width="100" />
    </el-table>
  </div>
</template>

<script setup>

  const tableData = [
    {
      ranking: 1,
      title: '全球领先的RWA一站式生态平台',
      click_num: 523,
      hot: 263,
      link: ''
    },
    {
      ranking: 2,
      title:
        "The world's leading one-stop RWA ecological platform",
      click_num: 337,
      hot: 176,
      link: ''
    },
    {
      ranking: 3,
      title:'共享行为价值共同体',
      click_num: 416,
      hot: 223,
      link: ''
    },
    {
      ranking: 4,
      title:
        'Shared behavioral value community',
      click_num: 292,
      hot: 145,
      link: ''
    },
    {
      ranking: 5,
      title: 'DeedMX全球领先的RWA',
      click_num: 173,
      hot: 110,
      link: ''
    }
  ]
</script>

<style scoped lang="scss"></style>
