-- ----------------------------
-- 背景图片地址
-- ----------------------------
SET @instance_background_img = 10;

-- ----------------------------
-- 版本
-- ----------------------------
SET @version = 1;

INSERT INTO `screen_instance` (
    `add_time`,
    `add_user`,
    `edit_time`,
    `edit_user`,
    `del_flag`,
    `instance_title`,
    `instance_background_color`,
    `instance_background_img`,
    `instance_height`,
    `instance_width`,
    `instance_version`,
    `instance_view_img`
)
VALUES
(
    '2020-01-30 16:46:14',
    '1',
    '1',
    '2020-01-30 16:46:14',
    '1',
    '#263546',
    @instance_background_img,
    '1080',
    '企业实时销售数据',
    '',
    '1920',
    @version
);

-- ----------------------------
-- \"database\":9
-- 更换为自己的ID
-- ----------------------------

SET @instance_id = @@IDENTITY;

-- ----------------------------
-- 左侧大地图
-- ----------------------------
INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_2_map\",\"name\":\"name\",\"value\":\"value\",\"legend\":\"legend\",\"chartType\":\"mapChina\"}',
    'mapChina',
    'false',
    '[]',
    '700',
    'chart0',
    '8000',
    '{\"title\":{\"text\":\"中国地图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"grid\":{\"top\":\"60\",\"left\":\"1%\",\"right\":\"4%\",\"bottom\":\"10%\",\"containLabel\":true},\"tooltip\":{\"trigger\":\"item\"},\"legend\":{\"orient\":\"vertical\",\"left\":\"left\",\"data\":[\"人数\"],\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"visualMap\":{\"min\":5000,\"max\":10000,\"left\":\"left\",\"top\":\"bottom\",\"text\":[\"高\",\"低\"],\"calculable\":true,\"color\":[\"#711D8E\",\"#484CDC\",\"#24CBFF\"]},\"geo\":{\"map\":\"china\"},\"backgroundColor\":\"rgba(255,255,255,0)\",\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"series\":[{\"name\":\"人数\",\"type\":\"map\",\"mapType\":\"china\",\"label\":{\"normal\":{\"show\":true}},\"data\":[{\"name\":\"北京\",\"value\":\"7040\"},{\"name\":\"天津\",\"value\":\"9876\"},{\"name\":\"上海\",\"value\":\"8278\"},{\"name\":\"重庆\",\"value\":\"8906\"},{\"name\":\"河北\",\"value\":\"9716\"},{\"name\":\"河南\",\"value\":\"5617\"},{\"name\":\"云南\",\"value\":\"7151\"},{\"name\":\"辽宁\",\"value\":\"6743\"},{\"name\":\"黑龙江\",\"value\":\"7202\"},{\"name\":\"湖南\",\"value\":\"9389\"},{\"name\":\"安徽\",\"value\":\"9485\"},{\"name\":\"山东\",\"value\":\"6896\"},{\"name\":\"新疆\",\"value\":\"7260\"},{\"name\":\"江苏\",\"value\":\"9821\"},{\"name\":\"浙江\",\"value\":\"7515\"},{\"name\":\"江西\",\"value\":\"5478\"},{\"name\":\"湖北\",\"value\":\"9371\"},{\"name\":\"广西\",\"value\":\"9802\"},{\"name\":\"甘肃\",\"value\":\"7963\"},{\"name\":\"山西\",\"value\":\"7768\"},{\"name\":\"内蒙古\",\"value\":\"6326\"},{\"name\":\"陕西\",\"value\":\"8913\"},{\"name\":\"吉林\",\"value\":\"8976\"},{\"name\":\"福建\",\"value\":\"6401\"},{\"name\":\"贵州\",\"value\":\"7137\"},{\"name\":\"广东\",\"value\":\"7612\"},{\"name\":\"青海\",\"value\":\"7889\"},{\"name\":\"西藏\",\"value\":\"8218\"},{\"name\":\"四川\",\"value\":\"5466\"},{\"name\":\"宁夏\",\"value\":\"7034\"},{\"name\":\"海南\",\"value\":\"5392\"},{\"name\":\"台湾\",\"value\":\"6428\"},{\"name\":\"香港\",\"value\":\"6802\"},{\"name\":\"澳门\",\"value\":\"6657\"}]}]}',
    'false',
    '920',
    '75',
    '305',
    @version
);

-- ----------------------------
-- 标题相关（共8个）
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '45',
    'chart1',
    '8000',
    '{\"title\":\"企业实时销售数据\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":35,\"fontWeight\":\"normal\",\"textAlign\":\"center\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '285',
    '50',
    '30',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '35',
    'chart2',
    '8000',
    '{\"title\":\"销售总额\",\"fontColor\":\"rgba(115, 170, 229, 1)\",\"fontSize\":30,\"fontWeight\":\"normal\",\"textAlign\":\"center\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '125',
    '830',
    '185',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '30',
    'chart3',
    '8000',
    '{\"title\":\"同比增长\",\"fontColor\":\"rgba(115, 170, 229, 1)\",\"fontSize\":22,\"fontWeight\":\"normal\",\"textAlign\":\"center\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '95',
    '95',
    '245',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart4',
    '8000',
    '{\"title\":\"一季度销售进度\",\"fontColor\":\"rgba(115, 170, 229, 1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"center\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '155',
    '1045',
    '35',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart5',
    '8000',
    '{\"title\":\"品类销售额占比\",\"fontColor\":\"rgba(115, 170, 229, 1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"center\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '150',
    '1465',
    '35',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart6',
    '8000',
    '{\"title\":\"销售品牌排名\",\"fontColor\":\"rgba(115, 170, 229, 1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"center\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '125',
    '1045',
    '395',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart7',
    '8000',
    '{\"title\":\"城市销售额排名\",\"fontColor\":\"rgba(115, 170, 229, 1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"center\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '150',
    '1470',
    '395',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    @version
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart8',
    '8000',
    '{\"title\":\"品牌分布\",\"fontColor\":\"rgba(115, 170, 229, 1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"center\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '95',
    '1045',
    '725',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart9',
    '8000',
    '{\"title\":\"核心城市销售增长额\",\"fontColor\":\"rgba(115, 170, 229, 1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"center\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '185',
    '1470',
    '685',
    @version
);

-- ----------------------------
-- 一季度销售进度
-- ----------------------------
INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_2_sp_pie\",\"value\":\"value\",\"chartType\":\"piePercent\"}',
    'piePercent',
    'false',
    '[]',
    '205',
    'chart10',
    '8000',
    '{\"title\":{\"text\":\"85%\",\"x\":\"center\",\"y\":\"center\",\"textStyle\":{\"color\":\"#98a0c4\",\"fontSize\":20}},\"series\":[{\"name\":\"main\",\"type\":\"pie\",\"radius\":[\"60%\",\"75%\"],\"center\":[\"50%\",\"50%\"],\"label\":{\"normal\":{\"show\":false}},\"data\":[{\"value\":85,\"itemStyle\":{\"normal\":{\"color\":\"rgba(119,96,246,1)\"}}},{\"value\":15,\"itemStyle\":{\"normal\":{\"color\":\"rgba(44,52,92,1)\"}}}]}],\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '410',
    '1040',
    '65',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select data from template_2_sv_count\",\"data\":\"data\",\"chartType\":\"counter\"}',
    'counter',
    'false',
    '[]',
    '45',
    'chart11',
    '8000',
    '{\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":25,\"fontWeight\":\"normal\",\"textAlign\":\"center\",\"backgroundColor\":\"rgba(255,255,255,0)\",\"thousandth\":true,\"prefix\":{\"title\":\"一季度销售额\",\"fontColor\":\"rgba(115, 170, 229, 1)\",\"fontSize\":18,\"fontWeight\":\"normal\"},\"unit\":{\"title\":\"万\",\"fontColor\":\"rgba(115, 170, 229, 1)\",\"fontSize\":18,\"fontWeight\":\"normal\"}}',
    'false',
    '180',
    '1140',
    '270',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart12',
    '8000',
    '{\"title\":\"同比增长\",\"fontColor\":\"rgba(115, 170, 229, 1)\",\"fontSize\":18,\"fontWeight\":\"normal\",\"textAlign\":\"center\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '80',
    '1140',
    '315',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select data from template_2_svg_count\",\"data\":\"data\",\"chartType\":\"counter\"}',
    'counter',
    'false',
    '[]',
    '30',
    'chart13',
    '8000',
    '{\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":25,\"fontWeight\":\"normal\",\"textAlign\":\"center\",\"backgroundColor\":\"rgba(255,255,255,0)\",\"thousandth\":true,\"prefix\":{\"title\":\"+\",\"fontColor\":\"rgba(231, 17, 17, 1)\",\"fontSize\":28,\"fontWeight\":\"normal\"},\"unit\":{\"title\":\"%\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":25,\"fontWeight\":\"normal\"}}',
    'false',
    '110',
    '1225',
    '310',
    @version
);
-- ----------------------------
-- 品牌销售额占比
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select name,value from template_2_ss_pie\",\"name\":\"name\",\"value\":\"value\",\"chartType\":\"pie2D\"}',
    'pie2D',
    'false',
    '[]',
    '315',
    'chart14',
    '8000',
    '{\"lunbo\":\"0\",\"title\":{\"text\":\"2D饼图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"color\":[\"#24CBFF\",\"#7760F6\",\"#484CDC\",\"#711D8E\"],\"tooltip\":{\"trigger\":\"item\"},\"legend\":{\"show\":true,\"orient\":\"horizontal\",\"top\":265,\"left\":55,\"data\":[\"数码类\",\"日用品\",\"快消类\",\"食品类\"],\"textStyle\":{\"color\":\"#fff\"}},\"series\":[{\"name\":\"访问来源\",\"type\":\"pie\",\"radius\":\"55%\",\"center\":[\"50%\",\"50%\"],\"data\":[{\"name\":\"数码类\",\"value\":\"371570\"},{\"name\":\"日用品\",\"value\":\"296016\"},{\"name\":\"快消类\",\"value\":\"161953\"},{\"name\":\"食品类\",\"value\":\"125822\"}],\"roseType\":\"angle\",\"label\":{\"normal\":{\"show\":false,\"position\":\"outside\",\"textStyle\":{\"color\":\"#ffffff\"}}},\"labelLine\":{\"normal\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"},\"smooth\":0.2,\"length\":10,\"length2\":20}},\"itemStyle\":{\"normal\":{\"color\":\"\",\"shadowBlur\":200,\"shadowColor\":\"rgba(0,0,0,0)\"}},\"animationType\":\"scale\",\"animationEasing\":\"elasticOut\",\"avoidLabelOverlap\":true}],\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '410',
    '1465',
    '65',
    @version
);
-- ----------------------------
-- 销售品牌排名
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_2_sbr_histogram\",\"y\":\"y\",\"x\":\"x\",\"chartType\":\"histogramGradientHorizontal\"}',
    'histogramGradientHorizontal',
    'false',
    '{}',
    '285',
    'chart15',
    '8000',
    '{\"title\":{\"text\":\"水平柱状图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"color\":[\"#24CBFF\"],\"tooltip\":{\"trigger\":\"axis\",\"axisPointer\":{\"type\":\"shadow\"}},\"grid\":{\"left\":\"5%\",\"right\":\"5%\",\"bottom\":\"5%\",\"containLabel\":true,\"top\":\"5%\"},\"legend\":{\"show\":true,\"orient\":\"vertical\",\"data\":[],\"top\":\"0\",\"left\":\"0\",\"textStyle\":{\"color\":\"#ffffff\"}},\"xAxis\":{\"data\":[],\"name\":\"\",\"type\":\"value\",\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"}},\"axisTick\":{\"show\":false},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}}},\"yAxis\":{\"name\":\"\",\"type\":\"category\",\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"}},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"data\":[\"三星\",\"vivo\",\"oppo\",\"华为\",\"小米\",\"iPhone\"]},\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"series\":{\"name\":\"\",\"type\":\"bar\",\"barWidth\":\"40%\",\"data\":[641731,775032,802838,954719,1281971,1570755],\"itemStyle\":{\"normal\":{\"barBorderRadius\":[80,80,80,80]}}},\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '410',
    '1040',
    '430',
    @version
);
-- ----------------------------
-- 城市销售额排名
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_2_sr_histogram\",\"y\":\"y\",\"x\":\"x\",\"chartType\":\"histogramGradient\"}',
    'histogramGradient',
    'false',
    '{}',
    '235',
    'chart16',
    '8000',
    '{\"title\":{\"text\":\"渐变柱状图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"legend\":{\"show\":false},\"backgroundColor\":\"\",\"grid\":{\"left\":\"3%\",\"right\":\"5%\",\"bottom\":\"5%\",\"containLabel\":true,\"top\":\"10%\"},\"xAxis\":{\"data\":[\"北京\",\"上海\",\"杭州\",\"深圳\",\"广州\",\"连云港\"],\"name\":\"\",\"axisTick\":{\"show\":false},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLabel\":{\"show\":true,\"inside\":false,\"textStyle\":{\"color\":\"#fff\"},\"rotate\":0,\"interval\":0},\"z\":10},\"yAxis\":{\"name\":\"\",\"axisTick\":{\"show\":false},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"}}},\"dataZoom\":[{\"type\":\"inside\"}],\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"series\":[{\"type\":\"bar\",\"itemStyle\":{\"normal\":{\"color\":\"rgba(255,255,255,0)\"}},\"barWidth\":\"30%\",\"barGap\":\"-100%\",\"barCategoryGap\":\"0\",\"data\":[[1339783],[1339783],[1339783],[1339783],[1339783],[1339783]],\"animation\":false},{\"type\":\"bar\",\"barWidth\":\"35%\",\"itemStyle\":{\"normal\":{\"color\":{\"x\":0,\"y\":0,\"x2\":0,\"y2\":1,\"type\":\"linear\",\"global\":false,\"colorStops\":[{\"offset\":0,\"color\":\"#24CBFF\"},{\"offset\":0.5,\"color\":\"#24CBFF\"},{\"offset\":1,\"color\":\"#24CBFF\"}]},\"barBorderRadius\":[80,80,80,80]},\"emphasis\":{\"color\":{\"x\":0,\"y\":0,\"x2\":0,\"y2\":1,\"type\":\"linear\",\"global\":false,\"colorStops\":[{\"offset\":0,\"color\":\"#24CBFF\"},{\"offset\":0.7,\"color\":\"#24CBFF\"},{\"offset\":1,\"color\":\"#24CBFF\"}]}}},\"data\":[1339783,1394842,786996,550908,612633,668681]}]}',
    'false',
    '410',
    '1465',
    '430',
    @version
);
-- ----------------------------
-- 品牌分布
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_2_bd_histogram\",\"y\":\"y\",\"x\":\"x\",\"chartType\":\"histogramGradient\"}',
    'histogramGradient',
    'false',
    '{}',
    '275',
    'chart17',
    '8000',
    '{\"title\":{\"text\":\"渐变柱状图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"legend\":{\"show\":false},\"backgroundColor\":\"\",\"grid\":{\"left\":\"10%\",\"right\":\"10%\",\"bottom\":\"5%\",\"containLabel\":true,\"top\":\"15%\"},\"xAxis\":{\"data\":[\"苹果\",\"三星\",\"华为\",\"oppo\",\"vivo\",\"小米\"],\"name\":\"\",\"axisTick\":{\"show\":false},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLabel\":{\"show\":true,\"inside\":false,\"textStyle\":{\"color\":\"#fff\"},\"rotate\":0,\"interval\":0},\"z\":10},\"yAxis\":{\"name\":\"\",\"axisTick\":{\"show\":false},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLabel\":{\"show\":false,\"textStyle\":{\"color\":\"#fff\"}}},\"dataZoom\":[{\"type\":\"inside\"}],\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"color\":[\"#24CBFF\"],\"series\":[{\"type\":\"pictorialBar\",\"itemStyle\":{\"normal\":{\"color\":\"rgba(255,255,255,0)\"}},\"barWidth\":\"30%\",\"barGap\":\"-100%\",\"barCategoryGap\":\"0\",\"data\":[[9102],[9102],[9102],[9102],[9102],[9102]],\"animation\":false},{\"type\":\"pictorialBar\",\"barCategoryGap\":\"-130%\",\"symbol\":\"path://M0,10 L10,10 C5.5,10 5.5,5 5,0 C4.5,5 4.5,10 0,10 z\",\"itemStyle\":{\"normal\":{\"opacity\":\"0.7\"},\"emphasis\":{\"itemStyle\":{\"opacity\":\"1\"}}},\"data\":[9896,8663,7458,9102,5114,8722]}]}',
    'false',
    '410',
    '1040',
    '755',
    @version
);
-- ----------------------------
-- 核心城市销售增长额
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_2_sg_line\",\"x\":\"x\",\"y\":\"y\",\"legend\":\"legend\",\"chartType\":\"lineStacking\"}',
    'lineStacking',
    'false',
    '[]',
    '310',
    'chart18',
    '8000',
    '{\"title\":{\"text\":\"堆叠折线图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"grid\":{\"top\":\"10%\",\"left\":\"5%\",\"right\":\"5%\",\"bottom\":\"5%\",\"containLabel\":true},\"legend\":{\"left\":\"1%\",\"top\":\"5%\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"tooltip\":{},\"dataset\":{\"source\":[[\"product\",\"北京\",\"上海\"],[\"2014\",\"883192\",\"710627\"],[\"2015\",\"643885\",\"358548\"],[\"2016\",\"658223\",\"707107\"],[\"2017\",\"521328\",\"316168\"],[\"2018\",\"833451\",\"485196\"]]},\"xAxis\":{\"name\":\"\",\"type\":\"category\",\"boundaryGap\":true,\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"},\"rotate\":0},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"yAxis\":{\"name\":\"\",\"type\":\"value\",\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":false,\"textStyle\":{\"color\":\"#fff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitArea\":{\"show\":false,\"areaStyle\":{\"color\":[\"rgba(250,250,250,0.3)\",\"rgba(200,200,200,0.3)\"]}},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"color\":[\"#24CBFF\",\"#7760F6\"],\"series\":[{\"type\":\"line\"},{\"type\":\"line\"}],\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '405',
    '1465',
    '720',
    @version
);
-- ----------------------------
-- 完善，修补
-- ----------------------------
