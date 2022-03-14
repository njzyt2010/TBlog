<template>
    <el-container>
        <el-header>
            <div>
                <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" />
                <div>asdfadf</div>
            </div>
            <div>afaddf</div>
        </el-header>
        <el-main>
            <div class="main-div">
                <!-- <Nuxt /> -->
            </div>
            <el-footer class="div-footer"
                       height="32px">
                copyright&copy;
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
    display: flex;
    // padding: 0 10px;
    & > div {
        cursor: pointer;
        padding: 10px;
        display: flex;
        line-height: 40px;

        & > div {
            padding: 0 5px;
        }
    }
}
.el-main {
    padding: 0 10px;
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