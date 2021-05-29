import React, { useState, useEffect } from 'react'
import Head from 'next/head'
import styles from '../public/style/pages/index.module.css'
import Header from '../components/Header'
import servicePath from '../config/apiUrl'

import { Row, Col, List, Icon, Pagination } from 'antd'
import Link from 'next/link'
import axios from 'axios'
export default function Home() {
  const [page, setPage] = useState(1)
  const [length, setLength] = useState(0)
  const [list, setList] = useState([])
  useEffect(() => {
    getList()
  }, [])
  //获取新闻列表
  const getList = (page = 1) => {
    axios({
      method: 'get',
      url: servicePath.getArticleList + page,
      withCredentials: false
    }).then(
      res => {
        setList(res.data.data)
        setLength(res.data.len)
      }
    )
  }
  //分页处理
  function onPageChange(page) {
    setPage(page)
    window.location.hash = `#${page}`;
    getList(page)
  }
  return (
    <div>
      <Head>
      </Head>
      <Header />
      <Row className="commmain" type="flex" justify="center">
        <Col className="commleft" xs={24} sm={24} md={15} lg={15} xl={15}>
          <List
            header={<div>新闻目录</div>}
            itemLayout="vertical"
            dataSource={list}
            renderItem={item => {
              return <List.Item>
                <div className={styles.listtitle}>
                  <Link href={{ pathname: '/detailed', query: { ID: item.ID } }}>
                    <a>{item.Title}</a>
                  </Link></div>
                <div className={styles.listicon}>
                  <span><Icon type="calendar" /> {item.UpdatedAt.split('T')[0] + ' ' + item.UpdatedAt.split('T')[1].split('+')[0]} </span>
                  <span>{item.Author}</span>
                </div>
              </List.Item>
            }}
          />
        </Col>
      </Row>

      <Row className="commmain" type="flex" justify="center">
        <Pagination defaultCurrent={page} current={page} onChange={onPageChange} total={length} pageSize={10}
        />
      </Row>
    </div>
  )
}


