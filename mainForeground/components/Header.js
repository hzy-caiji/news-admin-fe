import styles from '../public/style/components/header.module.css'
import { Row, Col, Icon, Menu } from 'antd'

import React, { memo } from 'react'
import Link from 'next/link'


const Header = () => {
    return (
        <div className={styles.header}>
            <Row type="flex" justify="center">
                <Col xs={21} sm={20} md={10} lg={15} xl={12}>
                    <span className={styles.headerlogo}><Link href={{ pathname: '/' }}>
                        <a> 今日新闻</a>
                    </Link></span>
                    <span className={styles.headertxt} >包含海量资讯的新闻服务平台,真实反映每时每刻的新闻热点</span>
                </Col>
            </Row>
        </div >)
}
export default memo(Header)