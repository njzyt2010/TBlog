<template>
    <el-container>
        <el-header>
            <div>
                <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" />
            </div>
            
        </el-header>
        <el-main>
            <div class="main-div">
                <!-- <Nuxt /> -->
            </div>
            <el-footer class="div-footer">
                copyright
            </el-footer>
        </el-main>
    </el-container>
</template>

<script>
export default {
    name: 'Blog'
}
</script>
<style lang="scss" scoped>
.common-layout {
    background-color: #ffffffcc;
}
.el-header {
    background-color: aquamarine;
    // padding: 0 10px;
    div {
        height: 60px;
    }
}
.el-main {
    // padding: 0 10px;
    .main-div {
        min-height: calc(100vh - 92px);
    }
    .div-footer {
        height: 32px;
        line-height: 32px;
        text-align: center;
        // padding: 0 10px;
    }
}
</style>