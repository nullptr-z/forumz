import { Fragment, useEffect, useState } from 'react'
import { Dialog, Transition } from '@headlessui/react'
import { getUserInfo } from '@/utils'
import axios from '@/utils/axios';

export default function Overlay(props) {
  const [open, setOpen] = useState(false)
  const { userId } = getUserInfo()

  useEffect(() => {
    setOpen(props.visible)
  }, [props.visible])

  const onClose = (v) => {
    setOpen(v)
    props?.onClose(v)
  }

  const onConfirm = async (e) => {
    const data = {
      userId: e.currentTarget.userId.value,
      targetId: e.currentTarget.targetId.value
    }
    console.log("【 data 】==>", data);
    let rsp = await axios.postForm('/user/addFriend', data)
    console.log("【 rsp 】==>", rsp);
  }

  return (
    <Transition.Root show={open} as={Fragment}>
      <Dialog as="div" className="relative z-10" onClose={onClose}>
        <Transition.Child
          as={Fragment}
          enter="ease-in-out duration-500"
          enterFrom="opacity-0"
          enterTo="opacity-100"
          leave="ease-in-out duration-500"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        >
          <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" >
            left
          </div>
        </Transition.Child>

        <div className="fixed inset-0 overflow-hidden">
          <div className="absolute inset-0 overflow-hidden">
            <div className="pointer-events-none fixed inset-y-0 right-0 flex max-w-full pl-10">
              <Transition.Child
                as={Fragment}
                enter="transform transition ease-in-out duration-500 sm:duration-700"
                enterFrom="translate-x-full"
                enterTo="translate-x-0"
                leave="transform transition ease-in-out duration-500 sm:duration-700"
                leaveFrom="translate-x-0"
                leaveTo="translate-x-full"
              >
                <Dialog.Panel className="pointer-events-auto relative w-screen max-w-md">
                  <Transition.Child
                    as={Fragment}
                    enter="ease-in-out duration-500"
                    enterFrom="opacity-0"
                    enterTo="opacity-100"
                    leave="ease-in-out duration-500"
                    leaveFrom="opacity-100"
                    leaveTo="opacity-0"
                  >
                    <div className="absolute left-0 top-0 -ml-8 flex pr-2 pt-4 sm:-ml-10 sm:pr-4">
                      <button
                        type="button"
                        className="relative rounded-md border-0 p-1 text-gray-300 mr-2 hover:text-white focus:outline-none focus:ring-2 focus:ring-white"
                        onClick={() => setOpen(false)}
                      >
                        <p>关闭</p>
                        {/* <span className="absolute -inset-2.5" /> */}
                        {/* <span className="sr-only">Close panel</span> */}
                      </button>
                    </div>
                  </Transition.Child>
                  <div className="flex h-full flex-col overflow-y-scroll bg-white py-6 shadow-xl">
                    <form onSubmit={onConfirm} action="#" method='POST' encType='application/x-www-form-unlencoded'>
                      <div className="px-4 sm:px-6">
                        <Dialog.Title className="text-base font-semibold leading-6 text-gray-900">
                          添加好友
                        </Dialog.Title>
                      </div>
                      <input name="userId" value={userId} className="hidden" />
                      <div className="relative mt-6 flex-1 px-4 sm:px-6">
                        <div htmlFor="targetId" className="flex relative w-full">
                          <input
                            id="targetId"
                            name="targetId"
                            type="text"
                            placeholder="对方的用户id"
                            className="flex w-full border rounded-xl focus:outline-none focus:border-indigo-300 pl-4 h-10"
                          />
                          <button
                            className="flex items-center justify-center bg-indigo-500 hover:bg-indigo-600 rounded-xl text-white px-4 py-1 flex-shrink-0"
                          >
                            添加好友
                          </button>
                        </div>
                      </div>
                    </form>
                  </div>
                </Dialog.Panel>
              </Transition.Child>
            </div>
          </div>
        </div>
      </Dialog>
    </Transition.Root >
  )
}
