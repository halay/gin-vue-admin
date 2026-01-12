<template>
  <div>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="name" label="插件标题" show-overflow-tooltip width="200">
        <template #default="{ row }">
          <a
            class="text-black dark:text-white decoration-black/20 dark:decoration-white/20 hover:text-active"
            target="_blank"
          >{{ row.name }}</a>
        </template>
      </el-table-column>
      <el-table-column prop="resume" label="简介" show-overflow-tooltip></el-table-column>
      <el-table-column prop="money" label="价格" width="100">
        <template #default="{ row }">
          <span v-if="row.money === 0">免费</span>
          <span v-else>￥{{ row.money }}</span>
        </template>
      </el-table-column>
    </el-table>
     <div class="gva-pagination">
      <el-pagination
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[5, 10, 20]"
        :total="total"
        layout="total, prev, pager, next"
        size="small"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
      />
    </div>
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
    }]
</script>

<style scoped lang="scss"></style>
