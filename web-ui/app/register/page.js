"use client"

import axios from '@/utils/axios';
import { useRouter } from 'next/navigation';
import React from 'react'


export default function Example() {
  const router = useRouter();
  console.log("【 router 】==>", router);

  const handleSubmit = async (event) => {
    event.preventDefault(); // 阻止表单默认提交行为

    const data = {
      // 使用 event.currentTarget 获取表单数据
      name: event.currentTarget.name.value,
      password: event.currentTarget.password.value,
      repassword: event.currentTarget.repassword.value,
    };
    let resp = await axios.postForm('/user/register', data, {})// withCredentials: true
    if (resp.status == 200 && resp.data.success) {
      setTimeout(() => {
        router.push('login')
      }, 500);
    }
    // 可以在此发送数据到服务器或进行其他处理
  };

  return (
    <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-sm">
        <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
          Register in to your account
        </h2>
      </div>

      <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
        <form className="space-y-6" onSubmit={handleSubmit} action="#" method="POST">
          <div>
            <label htmlFor="name" className="block text-sm font-medium leading-6 text-gray-900">
              name
            </label>
            <div className="mt-2">
              <input
                id="name"
                name="name"
                type="input"
                // autoComplete="name"
                required
                className="block w-full pl-2 rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
              />
            </div>
          </div>

          <div>
            <div className="flex items-center justify-between">
              <label htmlFor="password" className="block text-sm font-medium leading-6 text-gray-900">
                Password
              </label>
              <div className="text-sm">
                <a href="#" className="font-semibold text-indigo-600 hover:text-indigo-500">
                  Forgot password?
                </a>
              </div>
            </div>
            <div className="mt-2">
              <input
                id="password"
                name="password"
                type="password"
                autoComplete="current-password"
                required
                className="block w-full pl-2 rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
              />
            </div>
          </div>

          {/* 确认密码 */}
          <div>
            <div className="flex items-center justify-between">
              <label htmlFor="password" className="block text-sm font-medium leading-6 text-gray-900">
                Confirm Password
              </label>
              <div className="text-sm">
                <a href="#" className="font-semibold text-indigo-600 hover:text-indigo-500">
                  Confirm password?
                </a>
              </div>
            </div>
            <div className="mt-2">
              <input
                id="repassword"
                name="repassword"
                type="password"
                autoComplete="current-password"
                required
                className="block w-full pl-2 rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
              />
            </div>
          </div>

          <div>
            <button
              // type="submit"
              className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
            >
              Register in
            </button>
          </div>
        </form>

        <p className="mt-10 text-center text-sm text-gray-500">
          <a href="/login" className="font-semibold leading-6 text-indigo-600 hover:text-indigo-500">
            go to login
          </a>
        </p>
      </div>
    </div>
  )
}
