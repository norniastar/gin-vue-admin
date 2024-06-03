<template>
  <div>
    <!--     搜索区域-->
    <div class="gva-search-box">
      <el-form
        :inline="true"
        :model="searchInfo"
      >
        <el-form-item label="订单号">
          <el-input
            v-model="searchInfo.order_sn"
            placeholder="请输入订单号"
          />
        </el-form-item>
        <el-form-item label="订单类型">
          <el-select
            v-model="searchInfo.type"
            placeholder="请选择订单类型"
          >
            <el-option
              v-for="item in orderType"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="订单状态">
          <el-select
            v-model="searchInfo.status"
            placeholder="请选择订单类型"
          >
            <el-option
              v-for="item in orderStatus"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="支付类型">
          <el-select
            v-model="searchInfo.pay_type"
            placeholder="请选择订单类型"
          >
            <el-option
              v-for="item in payType"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="订单来源">
          <el-select
            v-model="searchInfo.source_type"
            placeholder="请选择订单类型"
          >
            <el-option
              v-for="item in sourceType"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="备注">
          <el-input
            v-model="searchInfo.remark"
            placeholder="请输入备注"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            icon="search"
            @click="onSubmit"
          >
            查询
          </el-button>
          <el-button
            icon="refresh"
            @click="onReset"
          >
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <!--    表格区域-->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          icon="delete"
          :disabled="!multipleSelection.length"
          @click="onDelete"
        >
          删除
        </el-button>
        <el-button
          type="primary"
          icon="plus"
          @click="addOrder"
        >
          新增订单
        </el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column
          align="left"
          type="selection"
          width="55"
        />
        <el-table-column
          align="left"
          label="订单号"
          prop="order_sn"
          width="120"
        />
        <el-table-column
          align="left"
          label="订单类型"
          prop="type"
          width="120"
        />
        <el-table-column
          align="left"
          label="状态"
          width="120"
        >
          <template #default="scope">
            <span>{{ getStatusLabel(scope.row.status) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="订单金额"
          prop="price"
          width="120"
        />
        <el-table-column
          align="left"
          label="支付金额"
          prop="pay_price"
          width="120"
        />
        <el-table-column
          align="left"
          label="支付类型"
          width="120"
        >
          <template #default="scope">
            <span>{{ getPayTypeLabel(scope.row.pay_type) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="订单来源"
          width="120"
        >
          <template #default="scope">
            <span>{{ getSourceTypeLabel(scope.row.source_type) }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="是否有效"
          width="120"
        >
          <template #default="scope">
            <el-switch
              v-model="scope.row.is_valid"
              inline-prompt
              :active-value="1"
              :inactive-value="2"
              @change="()=>{switchValid(scope.row)}"
            />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="备注"
          width="220"
        >
          <template #default="scope">
            <el-input
              v-if="scope.row.editing"
              v-model="scope.row.remark"
              placeholder="请输入备注"
              @change="addOrSaveOrder(scope.row)"
              @blur="scope.row.editing = false"
            />
            <span v-else @click="scope.row.editing = true">{{ scope.row.remark }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="操作"
        >
          <template #default="scope">
            <el-button
              icon="delete"
              type="primary"
              link
              @click="removeOrderFunc(scope.row)"
            >
              删除
            </el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              @click="openOrderEditFunc(scope.row)"
            >
              编辑
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <!--分页区域-->
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <!--    抽屉区域-->
    <el-drawer
      v-model="addOrderDialog"
      size="60%"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">订单</span>
          <div>
            <el-button @click="closeAddOrderDialog">
              取 消
            </el-button>
            <el-button
              type="primary"
              @click="enterAddOrderDialog"
            >
              确 定
            </el-button>
          </div>
        </div>
      </template>

      <el-form
        ref="orderForm"
        :rules="rules"
        :model="orderInfo"
        label-width="80px"
      >
        <el-form-item
          label="订单号"
          prop="order_sn"
        >
          <span v-if="dialogFlag === 'edit'">{{ orderInfo.order_sn }}</span>
          <span v-else>-</span>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>


<script setup>
// 导入外部包
import {
  getOrderList,
  addOrSaveOrder,
  removeOrder,
} from '@/api/orderList'
import {nextTick, ref} from 'vue'
import {ElMessage, ElMessageBox} from "element-plus";


// 变量定义
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({
  order_sn: '',
  type: '',
  status: '',
  pay_type: '',
  source_type: '',
  remark: ''
})

// 方法定义
// 重置
const onReset = () => {
  searchInfo.value = {}

  getTableData()
}
// 提交
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10

  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}
// 获取数据
const getTableData = async () => {
  const table = await getOrderList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const switchValid = async (row) => {
  // await nextTick()
  const res = await addOrSaveOrder(row);
  if (res.code !== 0) {
    ElMessage({
      type: 'error',
      message: '保存失败'
    });
  } else {
    ElMessage({
      type: 'success',
      message: '保存成功'
    });
  }
}

// 调用获取数据, 进入页面立即获取数据
getTableData()

// 多选事件
const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}
// 行内删除
const onDelete = async () => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const ids = []
    multipleSelection.value &&
    multipleSelection.value.forEach(item => {
      ids.push(item.ID)
    })
    const res = await removeOrder({ids})
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === ids.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}
// 多选删除
const removeOrderFunc = async (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const ids = []
    ids.push(row.ID)
    const res = await removeOrder({ids})
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

const openOrderEditFunc = (row) => {
  dialogFlag.value = 'edit'
  orderInfo.value = JSON.parse(JSON.stringify(row))
  console.log(orderInfo.value.order_sn)
  addOrderDialog.value = true
}

// 弹窗相关
const orderInfo = ref({
  ID: 0,
  order_sn:'',
  type: 0,
  status: 0,
  price: 0,
  pay_price: 0,
  pay_type: '',
  source_type: '',
  is_valid: 0,
  remark: ''
})

const rules = ref({
  type: [
    {required: true, message: '请选择订单类型', trigger: 'blur'}
  ]
})

const orderForm = ref(null)
const enterAddOrderDialog = async () => {

  orderForm.value.validate(async valid => {
    if (valid) {
      const req = {
        ...orderInfo.value
      }
      if (dialogFlag.value === 'add') {
        const res = await addOrSaveOrder(req)
        if (res.code === 0) {
          ElMessage({type: 'success', message: '创建成功'})
          await getTableData()
          closeAddOrderDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await addOrSaveOrder(req)
        if (res.code === 0) {
          ElMessage({type: 'success', message: '编辑成功'})
          await getTableData()
          closeAddOrderDialog()
        }
      }
    }
  })
}

const addOrderDialog = ref(false)
const closeAddOrderDialog = () => {
  orderForm.value.resetFields()
  addOrderDialog.value = false

}

const dialogFlag = ref('add')
const addOrder = () => {
  dialogFlag.value = 'add'
  addOrderDialog.value = true
}


const orderType = ref([
  {
    label: 'vip订单',
    value: 'vip'
  },
  {
    label: '时长订单',
    value: 'duration'
  }
])
const orderStatus = ref([
  {
    label: '待支付',
    value: '1'
  },
  {
    label: '已支付',
    value: '2'
  },
  {
    label: '已关闭',
    value: '3'
  },
  {
    label: '已退款',
    value: '4'
  }
])
const payType = ref([
  {
    label: '微信',
    value: 'wechatpay'
  },
  {
    label: '支付宝',
    value: 'alipay'
  }
])
const sourceType = ref([
  {
    label: '官网',
    value: 'official'
  },
  {
    label: '后台',
    value: 'backend'
  }
])
const getStatusLabel = (statusValue) => {
  switch (statusValue) {
    case 1:
      return '待支付';
    case 2:
      return '已支付';
    case 3:
      return '已关闭';
    case 4:
      return '已退款';
    default:
      return '未知状态';
  }
}
const getPayTypeLabel = (payTypeValue) => {
  switch (payTypeValue) {
    case 'wechatpay':
      return '微信';
    case 'alipay':
      return '支付宝';
    default:
      return '未知方式';
  }
}
const getSourceTypeLabel = (sourceType) => {
  switch (sourceType) {
    case 'official':
      return '官网';
    case 'backend':
      return '后台赠送';
    case 'offline':
      return '订单录入';
    default:
      return '未知方式';
  }
}

</script>

<style scoped lang="scss">

</style>