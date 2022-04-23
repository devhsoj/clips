import type {Clip} from './clip'

interface User {
    user_id:number
    username:string
    clips:Clip[]
    created_at:Date
}

async function getMe():Promise<User> {
    try {
        const res = await fetch('/api/user/me')

        if(res.status !== 200) {

            if(res.status === 401) {
                return Promise.resolve(null)
            }

            return Promise.reject('Failed to get user')
        }

        const user = <User> (await res.json())

        return Promise.resolve(user)
    } catch(err) {
        console.trace(err)

        return Promise.reject('Unexpected Error')
    }
}

export {
    User,
    getMe
}