# localid

Local ID will give you integers on demand that are not currently in use from the generator. This works by keeping an internal counter of the highest ID it's generated, storing "returned" IDs in a queue, and simply generating more as required.

This means that IDs may be reused from this generator, if you use `ReturnIDs()`.

The benefit of this approach compared to just constantly incrementing an integer is that your ID space stays relatively small, and works well for long-running applications that need to frequently have a cheap way to identify an object, but don't want to generate unnecessarily large IDs, or deal with potential integer overflows (assuming they return IDs frequently enough).

This is used in [Fossabot](https://fossabot.com) to uniquely identify tasks within the distributed scheduler processes.
