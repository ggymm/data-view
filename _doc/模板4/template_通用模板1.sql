-- ----------------------------
-- 背景图片ID
-- ----------------------------
SET @instance_background_img = 11;

-- ----------------------------
-- screen_instance（大屏实例）
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
    '2020-02-21 16:33:21',
    '1',
    '1',
    '2020-02-21 16:33:21',
    '1',
    '#263546',
    @instance_background_img,
    '1080',
    '通用模板1',
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
-- 左1柱状图
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
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_4_data1_histogram\",\"y\":\"y\",\"x\":\"x\",\"chartType\":\"histogramGradient\"}',
    'histogramGradient',
    'false',
    '{}',
    '200',
    'chart0',
    '8000',
    '{\"title\":{\"text\":\"渐变柱状图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"legend\":{\"show\":false},\"backgroundColor\":\"\",\"grid\":{\"left\":\"0\",\"right\":\"0\",\"bottom\":\"0\",\"containLabel\":true,\"top\":\"5%\"},\"xAxis\":{\"data\":[\"一月\",\"二月\",\"三月\",\"四月\",\"五月\"],\"name\":\"\",\"axisTick\":{\"show\":false},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLabel\":{\"show\":true,\"inside\":false,\"textStyle\":{\"color\":\"#fff\"},\"rotate\":0,\"interval\":0},\"z\":10},\"yAxis\":{\"name\":\"\",\"axisTick\":{\"show\":false},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"}}},\"dataZoom\":[{\"type\":\"inside\"}],\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"color\":[\"#00C8DC\"],\"series\":[{\"type\":\"bar\",\"itemStyle\":{\"normal\":{\"color\":\"rgba(255,255,255,0)\"}},\"barWidth\":\"30%\",\"barGap\":\"-100%\",\"barCategoryGap\":\"0\",\"data\":[[120],[120],[120],[120],[120]],\"animation\":false},{\"type\":\"bar\",\"barWidth\":\"40%\",\"itemStyle\":{\"normal\":{\"color\":{\"x\":0,\"y\":0,\"x2\":0,\"y2\":1,\"type\":\"linear\",\"global\":false,\"colorStops\":[{\"offset\":0,\"color\":\"#00C8DC\"},{\"offset\":0.5,\"color\":\"#00C8DC\"},{\"offset\":1,\"color\":\"#00C8DC\"}]}},\"emphasis\":{\"color\":{\"x\":0,\"y\":0,\"x2\":0,\"y2\":1,\"type\":\"linear\",\"global\":false,\"colorStops\":[{\"offset\":0,\"color\":\"#00C8DC\"},{\"offset\":0.7,\"color\":\"#00C8DC\"},{\"offset\":1,\"color\":\"#00C8DC\"}]}}},\"data\":[90,40,180,80,120]}]}',
    'false',
    '315',
    '80',
    '190',
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
    'chart1',
    '8000',
    '{\"title\":\"数据标题\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '105',
    '80',
    '120',
    @version
);
-- ----------------------------
-- 左2柱状图
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
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_4_data2_histogram\",\"y\":\"y\",\"x\":\"x\",\"chartType\":\"histogramGradientHorizontal\"}',
    'histogramGradientHorizontal',
    'false',
    '{}',
    '200',
    'chart2',
    '8000',
    '{\"title\":{\"text\":\"水平柱状图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"tooltip\":{\"trigger\":\"axis\",\"axisPointer\":{\"type\":\"shadow\"}},\"grid\":{\"left\":\"0\",\"right\":\"5%\",\"bottom\":\"0\",\"containLabel\":true,\"top\":\"0\"},\"legend\":{\"show\":true,\"orient\":\"vertical\",\"data\":[],\"top\":\"0\",\"left\":\"0\",\"textStyle\":{\"color\":\"#ffffff\"}},\"xAxis\":{\"data\":[],\"name\":\"\",\"type\":\"value\",\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"}},\"axisTick\":{\"show\":false},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}}},\"yAxis\":{\"name\":\"\",\"type\":\"category\",\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"}},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"data\":[\"一月\",\"二月\",\"三月\",\"四月\"]},\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"color\":[\"#5F5FFF\"],\"series\":{\"name\":\"\",\"type\":\"bar\",\"barWidth\":\"40%\",\"data\":[80,50,160,120]},\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '315',
    '80',
    '520',
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
    'chart3',
    '8000',
    '{\"title\":\"数据标题\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '105',
    '80',
    '465',
    @version
);
-- ----------------------------
-- 左3饼状图
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
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select name,value from template_4_data3_pie\",\"name\":\"name\",\"value\":\"value\",\"chartType\":\"pieRing\"}',
    'pieRing',
    'false',
    '[]',
    '160',
    'chart4',
    '8000',
    '{\"lunbo\":\"0\",\"title\":{\"text\":\"环形饼图\",\"left\":\"center\",\"top\":\"top\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"tooltip\":{\"trigger\":\"item\",\"formatter\":\"\"},\"color\":[\"#006EFF\",\"#5F5FFF\",\"#00C8DC\"],\"legend\":{\"show\":false,\"orient\":\"vertical\",\"top\":0,\"left\":0,\"data\":[\"500元以下\",\"1K~2K元\",\"0.5K~1K元\"],\"textStyle\":{\"color\":\"#fff\"}},\"series\":[{\"name\":\"\",\"type\":\"pie\",\"radius\":[\"70%\",\"100%\"],\"center\":[\"50%\",\"50%\"],\"avoidLabelOverlap\":false,\"label\":{\"normal\":{\"show\":false,\"position\":\"center\"},\"emphasis\":{\"show\":true,\"textStyle\":{\"fontSize\":\"15\",\"fontWeight\":\"bold\"},\"formatter\":\"\"}},\"labelLine\":{\"normal\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"},\"smooth\":0.2,\"length\":10,\"length2\":20}},\"data\":[{\"name\":\"500元以下\",\"value\":\"1800\"},{\"name\":\"1K~2K元\",\"value\":\"1200\"},{\"name\":\"0.5K~1K元\",\"value\":\"600\"}]}],\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '335',
    '70',
    '850',
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
    '{\"title\":\"数据标题\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '105',
    '80',
    '800',
    @version
);

-- ----------------------------
-- 中间大地图
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
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_4_data4_map\",\"name\":\"name\",\"value\":\"value\",\"legend\":\"legend\",\"chartType\":\"mapChina\"}',
    'mapChina',
    'false',
    '[]',
    '680',
    'chart6',
    '8000',
    '{\"title\":{\"text\":\"中国地图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"grid\":{\"top\":\"\",\"left\":\"\",\"right\":\"\",\"bottom\":\"\",\"containLabel\":true},\"tooltip\":{\"trigger\":\"item\"},\"legend\":{\"orient\":\"vertical\",\"left\":\"left\",\"data\":[\"thermal\"],\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"visualMap\":{\"show\":false,\"min\":50,\"max\":550,\"left\":\"left\",\"top\":\"bottom\",\"text\":[\"高\",\"低\"],\"color\":[\"#00C8DC\",\"#080B19\"],\"calculable\":true},\"geo\":{\"map\":\"china\",\"itemStyle\":{\"areaColor\":\"#323c48\",\"borderColor\":\"#111\"}},\"backgroundColor\":\"rgba(255,255,255,0)\",\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"series\":[{\"name\":\"thermal\",\"type\":\"map\",\"mapType\":\"china\",\"label\":{\"normal\":{\"show\":false,\"textStyle\":{\"color\":\"#ffffff\"}},\"emphasis\":{\"show\":false,\"textStyle\":{\"color\":\"#ffffff\"}}},\"itemStyle\":{\"normal\":{\"color\":\"transparent\",\"areaColor\":\"#080B19\",\"borderColor\":\"#00C8DC\"},\"emphasis\":{\"areaColor\":\"#1B2D4D\"}},\"data\":[{\"name\":\"四川\",\"value\":\"100\"},{\"name\":\"广东\",\"value\":\"200\"},{\"name\":\"香港\",\"value\":\"300\"},{\"name\":\"上海\",\"value\":\"250\"},{\"name\":\"青海\",\"value\":\"50\"},{\"name\":\"重庆\",\"value\":\"350\"},{\"name\":\"山东\",\"value\":\"450\"},{\"name\":\"辽宁\",\"value\":\"550\"}]}]}',
    'false',
    '955',
    '480',
    '105',
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
    '{\"title\":\"总销量\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '65',
    '505',
    '200',
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
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select data from template_4_data4_sum\",\"data\":\"data\",\"chartType\":\"counter\"}',
    'counter',
    'false',
    '[]',
    '45',
    'chart8',
    '8000',
    '{\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":40,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\",\"thousandth\":true,\"prefix\":{\"title\":\"\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":12,\"fontWeight\":\"normal\"},\"unit\":{\"title\":\"\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":12,\"fontWeight\":\"normal\"}}',
    'false',
    '235',
    '505',
    '140',
    @version
);
-- ----------------------------
-- 中间折线图
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
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select x,y,legend from template_4_data5_line\",\"x\":\"x\",\"y\":\"y\",\"legend\":\"legend\",\"chartType\":\"lineStackingArea\"}',
    'lineStackingArea',
    'false',
    '[]',
    '185',
    'chart9',
    '8000',
    '{\"title\":{\"text\":\"堆叠面积图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"grid\":{\"left\":\"0\",\"right\":\"5%\",\"bottom\":\"0\",\"containLabel\":true},\"tooltip\":{\"trigger\":\"axis\"},\"legend\":{\"show\":false,\"data\":[\"data\"],\"orient\":\"vertical\",\"left\":0,\"top\":0,\"textStyle\":{\"color\":\"#fff\"}},\"showSymbol\":\"true\",\"xAxis\":{\"name\":\"\",\"type\":\"category\",\"boundaryGap\":false,\"data\":[\"一月\",\"二月\",\"三月\",\"四月\",\"五月\",\"六月\",\"七月\"],\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"},\"rotate\":0},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"yAxis\":{\"name\":\"\",\"type\":\"value\",\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitArea\":{\"show\":false,\"areaStyle\":{\"color\":[\"rgba(250,250,250,0.3)\"]}},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"color\":[\"#00C8DC\"],\"series\":[{\"name\":\"data\",\"type\":\"line\",\"data\":[600,80,680,1200,600,400,800],\"areaStyle\":{\"color\":\"rgba(0,200,220,0.17)\"},\"smooth\":true,\"symbol\":\"none\"}],\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '895',
    '505',
    '815',
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
    'chart10',
    '8000',
    '{\"title\":\"数据标题\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '105',
    '505',
    '800',
    @version
);
-- ----------------------------
-- 右1折线图
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
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select x,y,legend from template_4_data6_line\",\"x\":\"x\",\"y\":\"y\",\"legend\":\"legend\",\"chartType\":\"lineStackingArea\"}',
    'lineStackingArea',
    'false',
    '[]',
    '205',
    'chart11',
    '8000',
    '{\"title\":{\"text\":\"堆叠面积图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"grid\":{\"top\":\"5%\",\"left\":\"0\",\"right\":\"5%\",\"bottom\":\"0\",\"containLabel\":true},\"tooltip\":{\"trigger\":\"axis\"},\"legend\":{\"show\":false,\"data\":[\"data1\",\"data2\"],\"orient\":\"vertical\",\"left\":0,\"top\":0,\"textStyle\":{\"color\":\"#fff\"}},\"xAxis\":{\"name\":\"\",\"type\":\"category\",\"boundaryGap\":false,\"data\":[\"一月\",\"二月\",\"三月\",\"四月\",\"五月\",\"六月\"],\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"},\"rotate\":0},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"yAxis\":{\"name\":\"\",\"type\":\"value\",\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitArea\":{\"show\":false,\"areaStyle\":{\"color\":[\"rgba(250,250,250,0.3)\",\"rgba(200,200,200,0.3)\"]}},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"color\":[\"#6E3CFF\",\"#36E3FF\"],\"series\":[{\"name\":\"data1\",\"type\":\"line\",\"symbol\":\"none\",\"data\":[100,450,300,1000,400,800],\"areaStyle\":{\"color\":\"rgba(109, 79, 250, 0.25)\"}},{\"name\":\"data2\",\"type\":\"line\",\"symbol\":\"none\",\"data\":[1100,800,1600,600,1000,700],\"areaStyle\":{\"color\":\"rgba(54, 227, 255, 0.15)\"}}],\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '320',
    '1520',
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
    '25',
    'chart12',
    '8000',
    '{\"title\":\"数据标题\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '105',
    '1520',
    '120',
    @version
);
-- ----------------------------
-- 右2柱状图
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
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_4_data7_histogram\",\"x\":\"x\",\"y\":\"y\",\"legend\":\"legend\",\"chartType\":\"histogramStacking\"}',
    'histogramStacking',
    'false',
    '{}',
    '228',
    'chart13',
    '8000',
    '{\"title\":{\"text\":\"堆叠柱状图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"grid\":{\"left\":\"0\",\"right\":\"5%\",\"bottom\":\"0\",\"containLabel\":true,\"top\":\"5%\"},\"legend\":{\"left\":\"1%\",\"top\":\"5%\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"tooltip\":{},\"dataset\":{\"source\":[[\"product\",\"data1\",\"data2\"],[\"一月\",\"80\",\"65\"],[\"二月\",\"50\",\"40\"],[\"三月\",\"160\",\"130\"],[\"四月\",\"120\",\"100\"],[\"五月\",\"40\",\"30\"]]},\"xAxis\":{\"name\":\"\",\"type\":\"category\",\"boundaryGap\":true,\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"},\"rotate\":0},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"yAxis\":{\"name\":\"\",\"type\":\"value\",\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitArea\":{\"show\":false,\"areaStyle\":{\"color\":[\"rgba(250,250,250,0.3)\",\"rgba(200,200,200,0.3)\"]}},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"color\":[\"#6E3CFF\",\"#36E3FF\"],\"series\":[{\"type\":\"bar\",\"itemStyle\":{\"color\":{\"x\":0,\"y\":0,\"x2\":0,\"y2\":1,\"type\":\"linear\",\"global\":false,\"colorStops\":[{\"offset\":0,\"color\":\"#6E3CFF\"},{\"offset\":1,\"color\":\"#6E3CFF\"}]}}},{\"type\":\"bar\",\"itemStyle\":{\"color\":{\"x\":0,\"y\":0,\"x2\":0,\"y2\":1,\"type\":\"linear\",\"global\":false,\"colorStops\":[{\"offset\":0,\"color\":\"#36E3FF\"},{\"offset\":1,\"color\":\"#36E3FF\"}]}}}],\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '323',
    '1515',
    '520',
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
    'chart14',
    '8000',
    '{\"title\":\"数据标题\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '105',
    '1515',
    '465',
    @version
);
-- ----------------------------
-- 右3折线图
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
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_4_data8_line\",\"x\":\"x\",\"y\":\"y\",\"legend\":\"legend\",\"chartType\":\"lineStacking\"}',
    'lineStacking',
    'false',
    '[]',
    '165',
    'chart15',
    '8000',
    '{\"title\":{\"text\":\"堆叠折线图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"grid\":{\"top\":\"5%\",\"left\":\"0\",\"right\":\"5%\",\"bottom\":\"0\",\"containLabel\":true},\"legend\":{\"left\":\"1%\",\"top\":\"5%\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"tooltip\":{},\"dataset\":{\"source\":[[\"product\",\"data\"],[\"一月\",\"600\"],[\"二月\",\"80\"],[\"三月\",\"680\"],[\"四月\",\"1200\"],[\"五月\",\"600\"],[\"六月\",\"400\"]]},\"xAxis\":{\"name\":\"\",\"type\":\"category\",\"boundaryGap\":true,\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"},\"rotate\":0},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"yAxis\":{\"name\":\"\",\"type\":\"value\",\"axisTick\":{\"show\":false},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitArea\":{\"show\":false,\"areaStyle\":{\"color\":[\"rgba(250,250,250,0.3)\",\"rgba(200,200,200,0.3)\"]}},\"axisLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#FFFFFF\"}}},\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"color\":[\"#5F5FFF\"],\"series\":[{\"type\":\"line\",\"smooth\":true,\"symbol\":\"none\"}],\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '320',
    '1515',
    '840',
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
    'chart16',
    '8000',
    '{\"title\":\"数据标题\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":20,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '105',
    '1515',
    '800',
    @version
);
-- ----------------------------
-- 大标题
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
    '50',
    'chart17',
    '8000',
    '{\"title\":\"产品2019年度销售\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":34,\"fontWeight\":\"normal\",\"textAlign\":\"left\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '300',
    '810',
    '25',
    @version
);
