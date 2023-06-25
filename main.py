import asyncio
import time
from bilibili_api import user


async def get_fans_and_pubs(uid: int) -> None:
    uid = uid
    u = user.User(uid=uid)
    r = await u.get_videos(pn=1)

    info = await u.get_relation_info()
    # debug print(info)
    follower_cnt = info["follower"]

    # video list
    video_list = r["list"]["vlist"]
    # debug print(video_list)

    current_time = time.time()
    seven_days_ago = current_time - (7 * 24 * 60 * 60)

    timestamp = int(seven_days_ago)
    pub_cnt = 0
    for v in video_list:
        if v["created"] >= timestamp:
            pub_cnt = pub_cnt + 1

    print(uid, ",", follower_cnt, ",", pub_cnt)


async def main() -> None:
    print("uid", ",fan count", ",pub count")
    uid = []
    for line in open("bili.csv"):
        if int(line) > 0:
            await get_fans_and_pubs(int(line))


if __name__ == '__main__':
    asyncio.get_event_loop().run_until_complete(main())

