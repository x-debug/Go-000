### 学习笔记

##### Hystrix Rolling window

```java
/**
 * Number of requests during rolling window.
 * Number that failed (failure + success + timeout + threadPoolRejected + semaphoreRejected).
 * Error percentage;
 */
public static class HealthCounts {
    private final long totalCount;
    private final long errorCount;
    private final int errorPercentage;

    HealthCounts(long total, long error) {
        this.totalCount = total;
        this.errorCount = error;
        if (totalCount > 0) {
            this.errorPercentage = (int) ((double) errorCount / totalCount * 100);
        } else {
            this.errorPercentage = 0;
        }
    }

    private static final HealthCounts EMPTY = new HealthCounts(0, 0);

    public long getTotalRequests() {
        return totalCount;
    }

    public long getErrorCount() {
        return errorCount;
    }

    public int getErrorPercentage() {
        return errorPercentage;
    }

    public HealthCounts plus(long[] eventTypeCounts) {
        long updatedTotalCount = totalCount;
        long updatedErrorCount = errorCount;

        long successCount = eventTypeCounts[HystrixEventType.SUCCESS.ordinal()];
        long failureCount = eventTypeCounts[HystrixEventType.FAILURE.ordinal()];
        long timeoutCount = eventTypeCounts[HystrixEventType.TIMEOUT.ordinal()];
        long threadPoolRejectedCount = eventTypeCounts[HystrixEventType.THREAD_POOL_REJECTED.ordinal()];
        long semaphoreRejectedCount = eventTypeCounts[HystrixEventType.SEMAPHORE_REJECTED.ordinal()];

        updatedTotalCount += (successCount + failureCount + timeoutCount + threadPoolRejectedCount + semaphoreRejectedCount);
        updatedErrorCount += (failureCount + timeoutCount + threadPoolRejectedCount + semaphoreRejectedCount);
        return new HealthCounts(updatedTotalCount, updatedErrorCount);
    }

    public static HealthCounts empty() {
        return EMPTY;
    }

    public String toString() {
        return "HealthCounts[" + errorCount + " / " + totalCount + " : " + getErrorPercentage() + "%]";
    }
}
```

##### 参考资料:

https://medium.com/@darek1024/hystrix-circuit-breaker-how-to-set-it-up-properly-84c75cfbe3ee

https://www.dazhuanlan.com/2019/10/02/5d94ba10dddae/

https://github.com/Netflix/Hystrix/wiki/How-it-Works#flow7

https://github.com/afex/hystrix-go